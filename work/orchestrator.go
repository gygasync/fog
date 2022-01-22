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
	RegisterQueue(name string)
}

type orchestrator struct {
	connection *amqp.Connection
	logger     common.Logger

	channel *amqp.Channel
	queues  map[string]amqp.Queue
}

func NewOrchestrator(connection string, logger common.Logger) *orchestrator {
	once.Do(func() {
		conn, err := amqp.Dial(connection)
		if err != nil {
			logger.Error("Failed to create orchestrator", err)
			panic(err)
		}

		ch, err := conn.Channel()
		if err != nil {
			logger.Fatal("Failed creating channel")
			panic(err)
		}

		instance = orchestrator{
			connection: conn,
			logger:     logger,
			channel:    ch,
			queues:     make(map[string]amqp.Queue),
		}
	})

	return &instance
}

func (o *orchestrator) PublishWork(work *definition.Work) {
	err := o.channel.Publish("",
		work.WorkType,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(work.Payload),
		})

	if err != nil {
		o.logger.Errorf("Error publishing work of type: %s", work.WorkType)
		return
	}

	o.logger.Infof("Published work of type: %s", work.WorkType)
}

func (o *orchestrator) StartResponseQeue(comms definition.Communication, response responseQueue) {
	o.logger.Infof("Started listening for: %s", response.responseType)
}

func (o *orchestrator) RegisterQueue(name string) {
	for item := range o.queues {
		if item == name {
			o.logger.Warnf("Queue: %s is already registered.", name)
			return
		}
	}

	q, err := o.channel.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		o.logger.Fatalf("Failed creating queue: %s", name)
		panic(err)
	}

	o.queues[name] = q
	o.logger.Infof("Queue: %s successfully created", name)
}
