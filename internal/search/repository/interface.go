package repository

import (
	"context"
	"zoomer/internal/models"
)

type SearchRepository interface {
	FindRoomBySearch(ctx context.Context, search *models.RoomSearch) ([]*models.Room, error)
}
