package queue

import (
    "github.com/streadway/amqp"
)

type Publisher struct {
    channel *amqp.Channel
}

func NewPublisher(conn *amqp.Connection) (*Publisher, error) {
    ch, err := conn.Channel()
	defer ch.Close()

    FailOnError(err, "Failed to open a channel")

    return &Publisher{
        channel: ch,
    }, nil
}

func (p *Publisher) Publish(queueName string, body []byte) error {
    err := p.channel.Publish(
        "",     // exchange
        queueName, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        body,
			DeliveryMode: amqp.Persistent,
        })
    return err
}

