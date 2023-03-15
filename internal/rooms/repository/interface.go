package repository

import (
	"context"
	"zoomer/internal/models"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *models.Room) error

	GetRoomsByUserId(ctx context.Context, userId string) ([]*models.Room, error)

	GetAllRooms(ctx context.Context) ([]*models.Room, error)

	CountRooms(ctx context.Context, userId string) (int, error)
}
