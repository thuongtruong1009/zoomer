package usecase

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/rooms/presenter"
)

type UseCase interface {
	CreateRoom(ctx context.Context, userId string, name string) error

	GetRoomsByUserId(ctx context.Context, userId string) ([]*models.Room, error)

	GetAllRooms(ctx context.Context) ([]*models.Room, error)

	VerifyContact(ctx context.Context, username string) bool

	//sync to redis
	GetChatHistory(ctx context.Context, username1, username2, fromTS, toTS string) *presenter.ChatResponse

	ContactList(ctx context.Context, username string) *presenter.ChatResponse
}
