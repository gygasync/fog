package work

import "fog/work/definition"

type IWorker interface {
	GetWorkType() string
	Work(work definition.Work) *definition.Response
}
