package tasks

import (
	"fog/common"

	"github.com/streadway/amqp"
)

type IWorkerGroup interface {
	PostTask(task ITask)
	GetDetails() []string
	Respond(response []byte)
}

type WorkerGroup struct {
	queue       amqp.Queue
	returnQueue amqp.Queue
	name        string
	channel     *amqp.Channel
	logger      common.Logger
	responder   IResponse
}

func NewWorkerGroup(name string, responder IResponse, connection *amqp.Connection, logger common.Logger) *WorkerGroup {
	ch, err := connection.Channel()
	failOnError(logger, err, "Failed to open channel")
	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(logger, err, "Failed to declare queue "+name)
	retQ, err := ch.QueueDeclare(
		"__return-"+name, // name
		false,            // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	failOnError(logger, err, "Failed to declare return queue "+name)
	return &WorkerGroup{logger: logger, queue: q, returnQueue: retQ, channel: ch, name: name}
}

func (w *WorkerGroup) PostTask(task ITask) {
	err := w.channel.Publish(
		"",
		w.name,
		false,
		false,
		amqp.Publishing{
			ContentType: task.GetType(),
			Body:        task.GetBytes(),
		})

	if err != nil {
		w.logger.Warn("Unable to post task ", err)
	}
}

func (w *WorkerGroup) Respond(response []byte) {
	err := w.channel.Publish(
		"",
		w.returnQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "response",
			Body:        response,
		})

	if err != nil {
		w.logger.Warn("Unable to post task ", err)
	}
}

func (w *WorkerGroup) Start() {
	w.logger.Infof("Started listening to worker group %s", w.name)
	msgs, err := w.channel.Consume(
		w.returnQueue.Name,
		w.returnQueue.Name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		w.logger.Fatalf("%s: %s", "Failed to consume channel ", err)
		return
	}
	w.logger.Infof("Started worker group return listener")
	forever := make(chan bool)

	go func() {
		for m := range msgs {
			w.responder.Handle(m.Body)
		}
	}()

	<-forever
}

func (w *WorkerGroup) GetDetails() []string {
	var result []string
	for i := 0; i < len(result); i++ {

	}
	return []string{"exif"}
}
