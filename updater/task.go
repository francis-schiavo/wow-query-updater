package updater

import (
	"fmt"
	"sync"
	"time"
)

type LogType int8

const (
	_ LogType = iota
	LtDebug
	LtInfo
	LtWarning
	LtError
)

type TaskLog struct {
	LogType LogType
	Message string
}

type RunnableTask interface {
	Run()
	GetName() string
	GetProgress() int
}

type Task struct {
	Name     string
	Progress int

	mux         sync.Mutex
	waitCond    *sync.Cond
	concurrency int
	delay       int
	waitGroup   sync.WaitGroup
	queue       chan int
	suspended   bool

	logChan *chan TaskLog

	manager     *TaskManager
}

func (task *Task) GetName() string {
	return task.Name
}

func (task *Task) log(logType LogType, message string, args ...interface{}) {
	if task.logChan == nil {
		return
	}

	*task.logChan <- TaskLog{
		LogType: logType,
		Message: fmt.Sprintf(message, args...),
	}
}

func (task *Task) resume() {
	task.manager.ResumeTimestamp = time.Now().Local().Add(time.Minute * 15)
	time.Sleep(60 * time.Minute)
	task.log(LtWarning, "RESUMING ALL WORKERS\n")
	task.suspended = false
	task.manager.Status = true
	task.waitCond.Broadcast()
}

func (task *Task) suspend(workerId int) {
	task.mux.Lock()
	defer task.mux.Unlock()

	if task.suspended {
		task.log(LtDebug, "[Worker %d]ALREADY SUSPENDED\n", workerId)
		return
	}

	task.log(LtWarning, "[Worker %d]SUSPENDING ALL WORKERS\n", workerId)
	task.suspended = true
	task.manager.Status = false

	go task.resume()
}
