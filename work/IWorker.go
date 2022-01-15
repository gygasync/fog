package work

type IWorker interface {
	GetWorkType() string
	Work(work workDefinition) *responseDefinition
}
