package work

import (
	"encoding/json"
	"fog/common"
	"fog/work/definition"
	"sync"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

var once sync.Once
var instance orchestrator

type IOrchestrator interface {
	PublishWork(work *definition.Work)
	StartWorker(worker IWorker)
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

func (o *orchestrator) StartWorker(worker IWorker) {
	msgs, err := o.channel.Consume(
		worker.GetWorkType(),
		uuid.NewString(),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		o.logger.Fatalf("Failed to consume queue: %s", worker.GetWorkType())
		panic(err)
	}
	o.logger.Infof("Started worker for: %s", worker.GetWorkType())

	forever := make(chan bool)

	go func() {
		for m := range msgs {
			msg := m.Body
			var workDefinition definition.Work
			err = json.Unmarshal(msg, &workDefinition)
			if err != nil {
				o.logger.Errorf("Error parsing work type: %s", worker.GetWorkType())
				return
			}
			o.logger.Infof("Start work | type: %s id: %s", worker.GetWorkType(), workDefinition.Id)
			response := worker.Work(workDefinition)
			o.logger.Infof("End work | type: %s id: %s | Time elapsed: %s", worker.GetWorkType(), workDefinition.Id, response.TimeCreated.Sub(workDefinition.TimeCreated))
		}
	}()

	<-forever

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
