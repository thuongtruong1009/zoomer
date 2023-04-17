package queue

import (
	"context"
	"github.com/streadway/amqp"
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
