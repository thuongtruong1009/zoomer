package hub

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
)

type IHub interface {
	CreateStream(ctx context.Context) string

	GetParticipants(ctx context.Context, roomID string) []*models.Participant

	InsertIntoStream(ctx context.Context, roomID string, client *models.Participant)

	DeleteStream(ctx context.Context, roomID string)

	Receiver(ctx context.Context, roomId string, client *models.Participant)

	Broadcaster()
}
