package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
	"zoomer/internal/auth"
	"zoomer/internal/models"
	"zoomer/internal/rooms/repository"
)

type roomUsecase struct {
	roomRepo rooms.RoomRepository
	userRepo auth.UserRepository
}

func NewRoomUseCase(roomRepo rooms.RoomRepository, userRepo auth.userRepository) rooms.Usecase {
	return &roomUsecase{roomRepo: roomRepo, userRepo: userRepo}
}

func (ru roomUsecase) CreateRoom(ctx context.Context, userId string, name string) error {
	room := &models.Room{
		Id:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		CreatedBy: userId,
	}

	count, err := ru.roomRepo.CountRoom(ctx, userId)
	if err != nil {
		return err
	}

	user, err := ru.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return err
	}

	if user.Limit > count {
		return ru.roomRepo.CreateRoom(ctx, room)
	}

	return errors.New("limit exceeded")

}

func (ru roomUsecase) GetRoomByUserId(ctx context.Context, userId string) ([]*models.Room, error) {
	rooms, err := ru.rooRepo.GetRoomByUserId(ctx, userId)

	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (ru roomUsecase) GetAllRooms(ctx context.Context) ([]*models.Room, error) {
	rooms, err := ru.roomRepo.GetAllRooms(ctx)

	if err != nil {
		return nil, err
	}
	return rooms, nil
}
