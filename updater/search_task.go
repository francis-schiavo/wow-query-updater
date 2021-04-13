package updater

import (
	"fmt"
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"reflect"
	"sync"
	"sync/atomic"
	"wow-query-updater/connections"
	"wow-query-updater/datasets"
)

type SearchTask struct {
	Task
	SearchMethod string
	ItemMethod   string
	ItemCallback ItemCallback

	totalItems     int32
	processedItems int32
}

func (task *SearchTask) GetProgress() int {
	if task.totalItems > 0 {
		return int(task.processedItems * 100 / task.totalItems)
	} else {
		return 0
	}
}

func (task *SearchTask) worker(workerId int) {
	endpointInterface := reflect.ValueOf(connections.WowClient).MethodByName(task.ItemMethod)

	for id := range task.queue {
		args := []reflect.Value{
			reflect.ValueOf(id),
			reflect.ValueOf((*blizzard_api.RequestOptions)(nil)),
		}
		task.log(LtDebug, "[Worker %d] Processing %s %d\n", workerId, task.Name, id)

		response := endpointInterface.Call(args)[0].Interface().(*blizzard_api.ApiResponse)

		task.log(LtDebug, "[Worker %d] Finished processing %s %d\n", workerId, task.Name, id)
		if response.Status == 200 {
			task.ItemCallback(response)
			task.log(LtDebug, "Updated %s %d successfully!\n", task.Name, id)
			atomic.AddInt32(&task.processedItems, 1)
			task.waitGroup.Done()
		} else if response.Status == 429 {


			// Insert the failed id into the queue to retry later
			task.queue <- id
			// Suspend all goroutines temporarily
			task.suspend(workerId)
		} else {
			task.manager.incFailedRequests()
			task.log(LtError, "Failed to update %s %d with status: %d\n", task.Name, id, response.Status)
			atomic.AddInt32(&task.processedItems, 1)
			task.waitGroup.Done()
		}

		// Wait for a while after a "too many requests" response
		if task.suspended {
			// Wait until it is resumed
			task.log(LtDebug, "[Worker %d] Waiting\n", workerId)
			task.waitCond.L.Lock()
			task.waitCond.Wait()
			task.waitCond.L.Unlock()
			task.log(LtDebug, "[Worker %d] Resumed\n", workerId)
		}
	}
	task.log(LtDebug, "[Worker %d] Exiting\n", workerId)
}

func (task *SearchTask) Run() {
	task.log(LtInfo, "Running task: %s\n", task.GetName())
	task.waitGroup = sync.WaitGroup{}
	task.queue = make(chan int, 130000)
	m := &sync.Mutex{}
	task.waitCond = sync.NewCond(m)
	task.suspended = false

	endpointInterface := reflect.ValueOf(connections.WowClient).MethodByName(task.SearchMethod)

	// Only start processing after we have all IDs
	for w := 1; w <= task.concurrency-1; w++ {
		go task.worker(w)
	}

	var jsonData datasets.SearchResult
	lastID := 0
	for {
		args := []reflect.Value{
			reflect.ValueOf(&blizzard_api.RequestOptions{
				QueryString: map[string]string {
					"orderby": "id",
					"id": fmt.Sprintf("[%d,]", lastID + 1),
					"_pageSize": "1000",
				},
			}),
		}
		response := endpointInterface.Call(args)[0].Interface().(*blizzard_api.ApiResponse)

		if response.Status != 200 {
			task.manager.incFailedRequests()
			task.log(LtError, "Failed to obtain search data with status: %d.\n", response.Status)
		}

		response.Parse(&jsonData)

		if len(jsonData.Results) == 0 {
			task.log(LtDebug, "No more data for task %s\n", task.Name)
			break
		}

		itemCount := len(jsonData.Results)
		atomic.AddInt32(&task.totalItems, int32(itemCount))
		task.waitGroup.Add(itemCount)

		for _, item := range jsonData.Results {
			lastID = item.Data.ID
			task.queue <- lastID
		}
	}

	task.waitGroup.Wait()
	task.log(LtInfo, "Task %s finished.\n", task.Name)
	close(task.queue)
}
