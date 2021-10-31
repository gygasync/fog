package tasks

import "fog/common"

type IOrchestrator interface {
	RegisterGroup(group IWorkerGroup) error
}

type Orchestrator struct {
	logger  common.Logger
	workers []IWorkerGroup
}

func NewOrchestrator(logger common.Logger) *Orchestrator {
	return &Orchestrator{logger: logger, workers: make([]IWorkerGroup, 0)}
}

func (o *Orchestrator) RegisterGroup(group IWorkerGroup) error {
	o.workers = append(o.workers, group)
	return nil
}
