package tasks

type IWork interface {
	Do()
}

type Work struct {
	parentTask ITask
	work       func(args ...interface{})
}

func NewWork(parentTask ITask, work func(args ...interface{})) *Work {
	return nil
}
