package queue

import (
	"context"
	"github.com/streadway/amqp"
	"log"
)

func RabbitMQAdapter(ctx context.Context, queueName string, body []byte) (<-chan IDelivery, error) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/zoomer")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	log.Println("-> Successfully connected to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	consumer := &IConsumer{}
	publisher := &IPublisher{}

	deliveries, err := consumer.Consumer(ctx, ch, queueName)
	FailOnError(err, "Failed to register a consumer")

	err = publisher.Publish(ctx, ch, queueName, body)
	FailOnError(err, "Failed to publish a message")

	return deliveries, nil
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s\n", msg, err)
	}
}
