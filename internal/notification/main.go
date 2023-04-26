package main

import (
	"context"
	"fmt"
	"zoomer/internal/notification/queue"
)

func main() {
	ctx := context.Background()
	deliveries, _ := queue.RabbitMQAdapter(ctx, "test", []byte("Hello World!"))

	for delivery := range deliveries {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Printf("Received a message: %s\n", string(delivery.Delivery.Body))
			delivery.Delivery.Ack(false)
		}
	}
}
