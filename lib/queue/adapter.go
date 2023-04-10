package queue

import (
	"log"
	"github.com/streadway/amqp"
)

func RabbitMQConnect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	log.Println("Successfully connected to RabbitMQ")
	return conn, nil
}