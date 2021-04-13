package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"reflect"
	"sync"
	"sync/atomic"
	"wow-query-updater/connections"
)

type ItemCallback func(data *blizzard_api.ApiResponse)

type IndexTask struct {
	Task
	IndexMethod     string
	IndexCollection string
	ItemMethod      string
	ItemCallback    ItemCallback

	totalItems     int32
	processedItems int32
}

func (task *IndexTask) GetProgress() int {
	if task.totalItems > 0 {
		return int(task.processedItems * 100 / task.totalItems)
	} else {
		return 0
	}
}

func (task *IndexTask) worker(workerId int) {
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

func (task *IndexTask) Run() {
	task.log(LtInfo, "Running task: %s\n", task.GetName())

	task.waitGroup = sync.WaitGroup{}
	task.queue = make(chan int)
	m := &sync.Mutex{}
	task.waitCond = sync.NewCond(m)
	task.suspended = false

	endpointInterface := reflect.ValueOf(connections.WowClient).MethodByName(task.IndexMethod)
	args := []reflect.Value{reflect.ValueOf((*blizzard_api.RequestOptions)(nil))}
	response := endpointInterface.Call(args)[0].Interface().(*blizzard_api.ApiResponse)

	if response.Status != 200 {
		task.manager.incFailedRequests()
		task.log(LtError, "Failed to obtain index with status: %d.\n", response.Status)
		return
	}

	for w := 1; w <= task.concurrency; w++ {
		go task.worker(w)
	}

	var jsonData map[string]interface{}
	response.Parse(&jsonData)

	itemCount := len(jsonData[task.IndexCollection].([]interface{}))
	task.totalItems = int32(itemCount)
	task.waitGroup.Add(itemCount)
	for _, item := range jsonData[task.IndexCollection].([]interface{}) {
		task.queue <- int(item.(map[string]interface{})["id"].(float64))
	}

	task.waitGroup.Wait()
	task.log(LtInfo, "Task %s finished.\n", task.Name)
	close(task.queue)
}
