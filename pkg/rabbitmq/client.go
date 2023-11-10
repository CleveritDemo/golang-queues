package rabbitmq

import (
	"fmt"
	"golang-queues/pkg/config/properties"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type EventHandler interface {
	HandleSubEvent(msg string)
}

func NewRabbitMQConn(cfg *properties.RabbitMQInfo) (*amqp.Connection, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)
	return amqp.Dial(connAddr)
}

func StartConsumer(queueName string, handler EventHandler, cfg *properties.RabbitMQInfo) {
	amqpConn, err := NewRabbitMQConn(cfg)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}
	defer amqpConn.Close()

	ch, err := amqpConn.Channel()
	if err != nil {
		log.Fatalf("Error getting channel: %v", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error declaring queue: %v", err)
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error consuming queue: %v", err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			handler.HandleSubEvent(string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
