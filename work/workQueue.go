package work

import "fog/work/definition"

type workQueue struct {
	workType     string
	orchestrator IOrchestrator
}

func NewWorkQueue(workType string, orchestrator IOrchestrator) *workQueue {
	return &workQueue{
		workType:     workType,
		orchestrator: orchestrator,
	}
}

func (w *workQueue) PostTask(task *interface{}) {
	data := definition.NewWorkDefinition(w.workType, task)
	w.orchestrator.PublishWork(data)
}
