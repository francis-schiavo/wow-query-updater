package updater

import (
	blizzard_api "github.com/francis-schiavo/blizzard-api-go"
	"reflect"
	"sync"
	"wow-query-updater/connections"
)

type RangeTask struct {
	Task

	RangeStart   int
	RangeEnd     int
	ItemMethod   string
	ItemCallback ItemCallback
}

func (task *RangeTask) worker(workerId int) {
	endpointInterface := reflect.ValueOf(connections.WowClient).MethodByName(task.ItemMethod)

	for id := range task.queue {
		args := []reflect.Value{
			reflect.ValueOf(id),
			reflect.ValueOf((*blizzard_api.RequestOptions)(nil)),
		}
		task.log(LT_DEBUG, "[Worker %d] Processing item %d\n", workerId, id)

		response := endpointInterface.Call(args)[0].Interface().(*blizzard_api.ApiResponse)
		if !response.Cached {
			task.rateLimiter <- 1
		}

		task.log(LT_DEBUG, "[Worker %d] Finished processing item %d\n", workerId, id)
		if response.Status == 200 {
			task.ItemCallback(response)
			task.log(LT_INFO, "Updated %s %d successfully!\n", task.Name, id)
			task.waitGroup.Done()
		} else if response.Status == 429 {
			// Insert the failed id into the queue to retry later
			task.queue <- id
			// Suspend all goroutines temporarily
			task.suspend(workerId)
		} else {
			task.log(LT_INFO, "Failed to update %s %d with status: %d\n", task.Name, id, response.Status)
			task.waitGroup.Done()
		}

		// Wait for a while after a "too many requests" response
		if task.suspended {
			// Wait until it is resumed
			task.log(LT_DEBUG, "[Worker %d] Waiting\n", workerId)
			task.waitCond.L.Lock()
			task.waitCond.Wait()
			task.waitCond.L.Unlock()
			task.log(LT_DEBUG, "[Worker %d] Resumed\n", workerId)
		}
	}
	task.log(LT_DEBUG, "[Worker %d] Exiting\n", workerId)
}

func (task *RangeTask) Run() {
	task.waitGroup = sync.WaitGroup{}
	task.queue = make(chan int)
	m := &sync.Mutex{}
	task.waitCond = sync.NewCond(m)
	task.suspended = false

	for w := 1; w <= task.concurrency; w++ {
		go task.worker(w)
	}

	task.rateLimiter = make(chan int, task.concurrency)
	go task.rateLimitWorker()

	for id := task.RangeStart; id <= task.RangeEnd; id++ {
		task.waitGroup.Add(1)
		task.queue <- id
	}

	task.waitGroup.Wait()
	task.log(LT_INFO, "Task %s finished.", task.Name)
	close(task.queue)
	close(task.rateLimiter)
}