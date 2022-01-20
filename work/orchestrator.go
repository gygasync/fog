package work

import "fog/work/definition"

type IOrchestrator interface {
	PublishWork(work *definition.Work)
	StartResponseQeue(comms definition.Communication, response responseQueue)
}
