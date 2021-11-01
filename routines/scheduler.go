package routines

type IWorker interface {
	Work(args ...interface{}) <-chan interface{}
}
