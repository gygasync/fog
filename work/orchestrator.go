package work

import "fog/work/definition"

type IOrchestrator interface {
	PublishWork(work *definition.Work)
	RecieveResponse(data []byte) *definition.Response
}
