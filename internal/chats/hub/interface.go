package hub

import (
	"context"
	"zoomer/internal/models"
)

type IHub interface {
	Receiver(ctx context.Context, client *models.Client)

	Broadcaster()
}
