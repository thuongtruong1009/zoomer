package repository

import (
	"context"
	"zoomer/internal/models"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *models.Room) error

	GetRoomByUserId(ctx context.Context, userId string) ([]*models.Room, error)

	GetAllRooms(ctx context.Context) ([]*models.Room, error)

	CountRooms(ctx context.Context) (int, error)
}
