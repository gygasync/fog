package tasks

type IWorkerGroup interface {
	RegisterWorker(worker IWorker) error
}

type WorkerGroup struct {
	workers  []IWorker
	capacity int
}

func NewWorkerGroup(capacity int) *WorkerGroup {
	return &WorkerGroup{workers: make([]IWorker, 0), capacity: capacity}
}

func (w *WorkerGroup) RegisterWorker(worker IWorker) error {
	w.workers = append(w.workers, worker)
	return nil
}
