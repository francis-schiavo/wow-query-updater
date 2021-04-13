package updater

type SimpleMethod func()

type SimpleTask struct {
	Name string
	Method SimpleMethod
}

func (task *SimpleTask) GetName() string {
	return task.Name
}

func (task *SimpleTask) Run() {
	task.Method()
}

func (task *SimpleTask) GetProgress() int {
	return 50
}