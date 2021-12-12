package tasks

import (
	"fog/common"

	"github.com/streadway/amqp"
)

type IWorkerGroup interface {
	PostTask(task ITask)
	GetDetails() []string
}

type WorkerGroup struct {
	queue       amqp.Queue
	returnQueue amqp.Queue
	name        string
	channel     *amqp.Channel
	logger      common.Logger
}

func NewWorkerGroup(name string, connection *amqp.Connection, logger common.Logger) *WorkerGroup {
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

func (w *WorkerGroup) GetDetails() []string {
	var result []string
	for i := 0; i < len(result); i++ {

	}
	return []string{"exif"}
}
