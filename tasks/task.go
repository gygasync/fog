package tasks

type Task struct {
	work   func(args ...interface{}) (interface{}, error)
	parent IWorker

	CurrentItem string
	State       string
}

type ITask interface {
}

func NewTask(parent IWorker, work func(args ...interface{}) (interface{}, error), args ...interface{}) *Task {
	return &Task{parent: parent, work: work}
}

func (t *Task) Do(args ...interface{}) error {
	go t.work(args...)
	return nil
}
