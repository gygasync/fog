package tasks

import (
	"fog/common"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type IWorker interface {
	NewTask(f func(args ...interface{}) (interface{}, error)) error
	Notify(identifier string) error
}

type Worker struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
	workerName string
	logger     common.Logger
	worker     func([]byte)
}

func NewWorker(connection *amqp.Connection, queueName string, worker func([]byte), logger common.Logger) *Worker {
	ch, err := connection.Channel()
	failOnError(logger, err, "Failed to open a channel")
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(logger, err, "Failed to declare a queue")
	return &Worker{connection: connection, channel: ch, queue: q, workerName: uuid.NewString(), logger: logger, worker: worker}

}

func (w *Worker) Start() {
	w.logger.Infof("Started worker %s", w.workerName)
	msgs, err := w.channel.Consume(
		w.queue.Name,
		w.workerName,
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
	w.logger.Infof("Started worker")
	forever := make(chan bool)

	go func() {
		for m := range msgs {
			w.worker(m.Body)
		}
	}()

	<-forever
}

// func (w *Worker) NewTask(f func(args ...interface{}) (interface{}, error)) error {
// 	task, err := NewTask(w, f)
// 	if err != nil {
// 		return err
// 	}

// 	go task.Do()
// }
