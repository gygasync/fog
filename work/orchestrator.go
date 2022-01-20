package work

import (
	"fog/common"
	"fog/work/definition"
	"sync"

	"github.com/streadway/amqp"
)

var once sync.Once
var instance orchestrator

type IOrchestrator interface {
	PublishWork(work *definition.Work)
	StartResponseQeue(comms definition.Communication, response responseQueue)
}

type orchestrator struct {
	connection *amqp.Connection
	logger     common.Logger
}

func NewOrchestrator(connection string, logger common.Logger) orchestrator {
	once.Do(func() {
		conn, err := amqp.Dial(connection)
		if err != nil {
			logger.Error("Failed to create orchestrator", err)
			panic(err)
		}
		instance = orchestrator{
			connection: conn,
			logger:     logger,
		}
	})

	return instance
}
