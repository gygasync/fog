package tasks

import (
	"fog/common"

	"github.com/streadway/amqp"
)

type IOrchestrator interface {
	RegisterGroup(group IWorkerGroup) error
}

type Orchestrator struct {
	connection *amqp.Connection
	logger     common.Logger
	workers    []IWorkerGroup
}

func NewOrchestrator(connection string, logger common.Logger) *Orchestrator {
	conn, err := amqp.Dial(connection)
	failOnError(logger, err, "Failed to connect to RabbitMq")
	return &Orchestrator{logger: logger, workers: make([]IWorkerGroup, 0), connection: conn}
}

func (o *Orchestrator) RegisterGroup(group IWorkerGroup) error {
	o.workers = append(o.workers, group)
	return nil
}

func (o *Orchestrator) GetConnection() *amqp.Connection {
	return o.connection
}

func failOnError(logger common.Logger, err error, msg string) {
	if err != nil {
		logger.Fatalf("%s: %s", msg, err)
	}
}
