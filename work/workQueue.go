package work

import "fog/work/definition"

type IWorkQueue interface {
	PostTask(task interface{})
}

type workQueue struct {
	workType     string
	orchestrator IOrchestrator
}

func NewWorkQueue(workType string, orchestrator IOrchestrator) *workQueue {
	orchestrator.RegisterQueue(workType)
	return &workQueue{
		workType:     workType,
		orchestrator: orchestrator,
	}
}

func (w *workQueue) PostTask(task interface{}) {
	data := definition.NewWorkDefinition(w.workType, task)
	w.orchestrator.PublishWork(data)
}
