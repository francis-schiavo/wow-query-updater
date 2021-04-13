package updater

import (
	"sync"
	"sync/atomic"
	"time"
)

type TaskManager struct {
	Concurrency         int
	delay               int
	LogLevel            LogType
	CurrentTask         RunnableTask
	Progress            int
	Status              bool

	taskList []RunnableTask

	taskStart    int64
	taskEnd      int64
	taskProgress int64
	logChannel   *chan TaskLog

	FailedRequests  int32
	ResumeTimestamp time.Time
}

func NewTaskManager(concurrency int, logLevel LogType) *TaskManager {
	return &TaskManager{
		Concurrency:    concurrency,
		LogLevel:       logLevel,
		Progress:       0,
		Status:         true,
		FailedRequests: 0,
	}
}

func (manager *TaskManager) incFailedRequests() {
	atomic.AddInt32(&manager.FailedRequests, 1)
}

func (manager *TaskManager) AddIndexTask(name string, indexMethod string, indexCollection string, itemMethod string, updateCallback ItemCallback) {
	task := &IndexTask{
		Task: Task{
			Name:        name,
			mux:         sync.Mutex{},
			concurrency: manager.Concurrency,
			delay:       manager.delay,

			logChan: manager.logChannel,
			manager: manager,
		},
		IndexMethod:     indexMethod,
		IndexCollection: indexCollection,
		ItemMethod:      itemMethod,
		ItemCallback:    updateCallback,
	}
	manager.taskList = append(manager.taskList, task)
}

func (manager *TaskManager) AddIndexTaskLimited(name string, indexMethod string, indexCollection string, itemMethod string, updateCallback ItemCallback, concurrency int) {
	task := &IndexTask{
		Task: Task{
			Name:        name,
			mux:         sync.Mutex{},
			concurrency: concurrency,
			delay:       manager.delay,

			logChan: manager.logChannel,
			manager: manager,
		},
		IndexMethod:     indexMethod,
		IndexCollection: indexCollection,
		ItemMethod:      itemMethod,
		ItemCallback:    updateCallback,
	}
	manager.taskList = append(manager.taskList, task)
}

func (manager *TaskManager) AddSearchTask(name string, indexMethod string, itemMethod string, updateCallback ItemCallback) {
	task := &SearchTask{
		Task: Task{
			Name:        name,
			mux:         sync.Mutex{},
			concurrency: manager.Concurrency,
			delay:       manager.delay,

			logChan: manager.logChannel,
			manager: manager,
		},
		SearchMethod: indexMethod,
		ItemMethod:   itemMethod,
		ItemCallback: updateCallback,
	}
	manager.taskList = append(manager.taskList, task)
}

func (manager *TaskManager) AddMediaTask(name string, indexModel interface{}, mediaMethod string, mediaCallback MediaCallback) {
	task := &MediaTask{
		Task: Task{
			Name:        name,
			mux:         sync.Mutex{},
			concurrency: manager.Concurrency,
			delay:       manager.delay,

			logChan: manager.logChannel,
			manager: manager,
		},
		IndexModel:    indexModel,
		MediaMethod:   mediaMethod,
		MediaCallback: mediaCallback,
	}
	manager.taskList = append(manager.taskList, task)
}

func (manager *TaskManager) AddSimpleTask(name string, method SimpleMethod) {
	task := &SimpleTask{
		Name:   name,
		Method: method,
	}
	manager.taskList = append(manager.taskList, task)
}

func (manager *TaskManager) Run(start string, only []string, doneChannel chan int) {
	started := start == ""
	maxProgress := len(manager.taskList)
	for progress, task := range manager.taskList {
		manager.Progress = progress * 100.0 / maxProgress

		if task.GetName() == start {
			started = true
		}

		if !started {
			continue
		}

		if len(only) > 0 && !containString(task.GetName(), only) {
			continue
		}

		manager.CurrentTask = task
		task.Run()
	}
	close(doneChannel)
}
