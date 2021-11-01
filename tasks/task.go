package tasks

type Task struct {
	work   func(args ...interface{}) <-chan interface{}
	parent IWorker
}

type ITask interface {
}

func NewTask(parent IWorker, work func(args ...interface{}) <-chan interface{}, args ...interface{}) *Task {
	return &Task{parent: parent, work: work}
}

func (t *Task) Do(args ...interface{}) <-chan interface{} {
	result := make(chan interface{})
	go func() {
		defer close(result)
		result <- t.work(args...)
	}()

	// t.parent.Notify()
	return result
}
