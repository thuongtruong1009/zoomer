package adapter

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type RmqConfig struct {
	URL      string
	WaitTime time.Duration
	Attempts int
}

type RmqConnection struct {
	ConsumerExchange string
	RmqConfig
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Delivery   <-chan amqp.Delivery
}

func New(consumerExchange string, cfg RmqConfig) *RmqConnection {
	return &RmqConnection{
		ConsumerExchange: consumerExchange,
		RmqConfig:        cfg,
	}
}

func (c *RmqConnection) AttemptsConnect() error {
	var err error
	for i := c.Attempts; i > 0; i-- {
		if err = c.connect(); err == nil {
			break
		}
		log.Printf("Failed to connect to RabbitMQ, attempts left:  %d", i)
		time.Sleep(c.WaitTime)
	}

	if err != nil {
		return fmt.Errorf("Failed when reconnecting to RabbitMQ: %w", err)
	}
	return nil
}

func (c *RmqConnection) connect() error {
	var err error
	c.Connection, err = amqp.Dial(c.URL)
	if err != nil {
		return fmt.Errorf("Failed to setup dial connect to RabbitMQ: %w", err)
	}

	err = c.Channel.ExchangeDeclare(
		c.ConsumerExchange,
		"fanout",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to declare an exchange: %w", err)
	}

	queue, err := c.Channel.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to declare a queue: %w", err)
	}

	err = c.Channel.QueueBind(
		queue.Name,
		"",
		c.ConsumerExchange,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("Failed to bind a queue: %w", err)
	}

	c.Delivery, err = c.Channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("Failed to register a consumer: %w", err)
	}

	return nil
}
