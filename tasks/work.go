package tasks

type IWork interface {
	Do(args ...interface{})
}

type Work struct {
	parentTask ITask
	work       func(args ...interface{})
}

func NewWork(parentTask ITask, work func(args ...interface{})) *Work {
	return nil
}

func (w *Work) Do(args ...interface{}) {
	go w.work(args...)
}
