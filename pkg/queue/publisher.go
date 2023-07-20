package queue

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

// func NewPublisher(conn *amqp.Connection) (*Publisher, error) {
//     ch, err := conn.Channel()
// 	defer ch.Close()

//     FailOnError(err, "Failed to open a channel")

//     return &Publisher{
//         channel: ch,
//     }, nil
// }

func (p *IPublisher) Publish(ctx context.Context, ch *amqp.Channel, queueName string, body []byte) error {
	err := ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:  "text/plain", //"application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
		})

	FailOnError(err, "Failed to publish a message")
	return nil
}
