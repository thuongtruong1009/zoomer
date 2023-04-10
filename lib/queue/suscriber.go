package queue

import (
	"log"
    "github.com/streadway/amqp"
)

type Consumer struct {
    channel  *amqp.Channel
    queue    amqp.Queue
    messages <-chan amqp.Delivery
}

func NewConsumer(conn *amqp.Connection, queueName string) (*Consumer, error) {
    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }
    q, err := ch.QueueDeclare(
        queueName, // name
        false,     // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    if err != nil {
        return nil, err
    }

	prefetchCount := 1 * 4 // 4 messages at a time
	err = ch.Qos(prefetchCount, 0, false)
	if err != nil {
		return nil, err
	}

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // arguments
    )
    if err != nil {
        return nil, err
    }

	log.Println("Waiting for messages")

    return &Consumer{
        channel:  ch,
        queue:    q,
        messages: msgs,
    }, nil
}

