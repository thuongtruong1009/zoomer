package hub

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
)

type IHub interface {
	Receiver(ctx context.Context, client *models.Client)

	Broadcaster()
}
