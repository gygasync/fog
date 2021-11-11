package tasks

import (
	"fog/common"

	"github.com/streadway/amqp"
)

type IWorkerGroup interface {
	PostTask(task ITask)
}

type WorkerGroup struct {
	queue    amqp.Queue
	name     string
	channel  *amqp.Channel
	exchange string
	logger   common.Logger
}

func NewWorkerGroup(name string, exchange string, connection *amqp.Connection, logger common.Logger) *WorkerGroup {
	ch, err := connection.Channel()
	failOnError(logger, err, "Failed to open channel")
	q, err := ch.QueueDeclare(
		exchange, // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(logger, err, "Failed to declare queue "+name)
	return &WorkerGroup{logger: logger, queue: q, channel: ch, exchange: exchange, name: name}
}

func (w *WorkerGroup) PostTask(task ITask) {
	err := w.channel.Publish(
		w.exchange,
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
