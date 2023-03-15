package usecase

import (
	"context"
	"zoomer/internal/models"
)

type UseCase interface {
	CreateRoom(ctx context.Context, userId string, name string) error

	GetRoomsByUserId(ctx context.Context, userId string) ([]*models.Room, error)

	GetAllRooms(ctx context.Context) ([]*models.Room, error)
}
