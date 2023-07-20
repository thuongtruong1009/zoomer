package repository

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
)

type SearchRepository interface {
	FindRoomBySearch(ctx context.Context, search *models.RoomSearch) ([]*models.Room, error)
}
