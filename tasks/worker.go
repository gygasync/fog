package tasks

type IWorker interface {
	NewTask(f func(args ...interface{}) (interface{}, error)) error
	Notify() error
}

type Worker struct {
	Name     string
	capacity int
	Queue    chan ITask
}

func NewWorker(name string, capacity int) *Worker {
	return &Worker{Name: name, capacity: capacity, Queue: make(chan ITask, capacity)}
}

// func (w *Worker) NewTask(f func(args ...interface{}) (interface{}, error)) error {
// 	task, err := NewTask(w, f)
// 	if err != nil {
// 		return err
// 	}

// 	go task.Do()
// }
