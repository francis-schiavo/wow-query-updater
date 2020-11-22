package updater

import (
	"fmt"
	"sync"
)

type TaskManager struct {
	Concurrency int
	Delay      int
	LogLevel LogType

	taskList []RunnableTask

	taskStart    int64
	taskEnd      int64
	taskProgress int64
	logChannel   chan TaskLog
}

func NewTaskManager(concurrency int, delay int, logLevel LogType) *TaskManager {
	return &TaskManager{
		Concurrency: concurrency,
		Delay:       delay,
		LogLevel: logLevel,

		taskStart:    0,
		taskEnd:      0,
		taskProgress: 0,
	}
}

func (manager *TaskManager) AddIndexTask(name string, indexMethod string, indexCollection string, itemMethod string, updateCallback ItemCallback) {
	task := &IndexTask{
		Task:          Task{
			Name:        name,
			mux:         sync.Mutex{},
			concurrency: manager.Concurrency,
			delay:       manager.Delay,

			start:       &manager.taskStart,
			end:         &manager.taskEnd,
			progress:    &manager.taskProgress,
			logChan:     &manager.logChannel,
		},
		IndexMethod: indexMethod,
		IndexCollection: indexCollection,
		ItemMethod: itemMethod,
		ItemCallback:  updateCallback,
	}
	manager.taskList = append(manager.taskList, task)
}

func (manager *TaskManager) AddIndexTaskLimited(name string, indexMethod string, indexCollection string, itemMethod string, updateCallback ItemCallback, concurrency int) {
	task := &IndexTask{
		Task:          Task{
			Name:        name,
			mux:         sync.Mutex{},
			concurrency: concurrency,
			delay:       manager.Delay,

			start:       &manager.taskStart,
			end:         &manager.taskEnd,
			progress:    &manager.taskProgress,
			logChan:     &manager.logChannel,
		},
		IndexMethod: indexMethod,
		IndexCollection: indexCollection,
		ItemMethod: itemMethod,
		ItemCallback:  updateCallback,
	}
	manager.taskList = append(manager.taskList, task)
}

func (manager *TaskManager) AddSearchTask(name string, indexMethod string, itemMethod string, updateCallback ItemCallback) {
	task := &SearchTask{
		Task:          Task{
			Name:        name,
			mux:         sync.Mutex{},
			concurrency: manager.Concurrency,
			delay:       manager.Delay,

			start:       &manager.taskStart,
			end:         &manager.taskEnd,
			progress:    &manager.taskProgress,
			logChan:     &manager.logChannel,
		},
		SearchMethod: indexMethod,
		ItemMethod: itemMethod,
		ItemCallback:  updateCallback,
	}
	manager.taskList = append(manager.taskList, task)
}

func (manager *TaskManager) AddRangeTask(name string, rangeStart int, rangeEnd int, itemMethod string, updateCallback ItemCallback) {
	task := &RangeTask{
		Task:          Task{
			Name:        name,
			mux:         sync.Mutex{},
			concurrency: manager.Concurrency,
			delay:       manager.Delay,

			start:       &manager.taskStart,
			end:         &manager.taskEnd,
			progress:    &manager.taskProgress,
			logChan:     &manager.logChannel,
		},
		RangeStart: rangeStart,
		RangeEnd: rangeEnd,
		ItemMethod: itemMethod,
		ItemCallback:  updateCallback,
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

func (manager *TaskManager) LogMonitor() {
	for log := range manager.logChannel {
		if log.LogType >= manager.LogLevel {
			print(log.Message)
		}
	}
}

func (manager *TaskManager) Run() {
	manager.logChannel = make(chan TaskLog)
	go manager.LogMonitor()

	for _, task := range manager.taskList {
		manager.logChannel <- TaskLog{
			LogType: LT_WARNING,
			Message: fmt.Sprintf("Running task: %s\n", task.GetName()),
		}
		task.Run()
	}
	close(manager.logChannel)
}