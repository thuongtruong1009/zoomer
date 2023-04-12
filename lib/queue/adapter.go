package queue

import (
	"log"
	"github.com/streadway/amqp"
)

func RabbitMQConnect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	FailOnError(err, "Failed to connect to RabbitMQ")
	log.Println("Successfully connected to RabbitMQ")

	return conn, nil
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s\n", msg, err)
	}
}
