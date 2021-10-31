package tasks

type IWorker interface {
	EnqueueTask(task ITask) error
	EnqueueTasks(tasks []ITask) (int, error)
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

func (w *Worker) EnqueueTask(task ITask) error {
	w.Queue <- task
	return nil
}

func (w *Worker) Notify() error {
	w.capacity--
	return nil
}
