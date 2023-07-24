package queue

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type IDelivery struct {
	Delivery amqp.Delivery
	Ctx      context.Context
}

type IPublisher struct {
	ctx    context.Context
	cancel context.CancelFunc
}

type IConsumer struct {
	ctx    context.Context
	cancel context.CancelFunc
}
