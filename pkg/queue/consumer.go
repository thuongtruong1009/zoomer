package queue

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func (c *IConsumer) Consumer(ctx context.Context, ch *amqp.Channel, queueName string) (<-chan IDelivery, error) {
	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)

	FailOnError(err, "Failed to register a consumer")

	log.Println("::: Waiting for messages")

	deliveries := make(chan IDelivery)
	go func() {
		for msg := range msgs {
			select {
			case <-ctx.Done():
				return
			default:
				delivery := IDelivery{
					Delivery: msg,
					Ctx:      ctx,
				}
				deliveries <- delivery
			}
		}
		close(deliveries)
	}()
	return deliveries, nil
}
