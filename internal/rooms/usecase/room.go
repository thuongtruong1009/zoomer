package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
	"fmt"
	"zoomer/internal/models"
	auth "zoomer/internal/auth/repository"
	"zoomer/internal/rooms/repository"
	"zoomer/internal/rooms/presenter"
)

type roomUsecase struct {
	roomRepo repository.RoomRepository
	userRepo auth.UserRepository
}

func NewRoomUseCase(roomRepo repository.RoomRepository, userRepo auth.UserRepository) UseCase {
	return &roomUsecase{
		roomRepo: roomRepo,
		userRepo: userRepo,
	}
}

func (ru roomUsecase) CreateRoom(ctx context.Context, userId string, name string) error {
	room := &models.Room{
		Id:        uuid.New().String(),
		Name:      name,
		CreatedAt: time.Now(),
		CreatedBy: userId,
	}

	count, err := ru.roomRepo.CountRooms(ctx, userId)
	if err != nil {
		return err
	}

	user, err := ru.userRepo.GetUserById(ctx, userId)
	if err != nil {
		return err
	}

	if user.Limit > count {
		return ru.roomRepo.CreateRoom(ctx, room)
	} else {
		return errors.New("limit exceeded")
	}

}

func (ru roomUsecase) GetRoomsByUserId(ctx context.Context, userId string) ([]*models.Room, error) {
	rooms, err := ru.roomRepo.GetRoomsByUserId(ctx, userId)
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

// sync to redis
func (ru roomUsecase) VerifyContact(ctx context.Context, username string) bool {
	_, err := ru.userRepo.GetUserByUsername(context.Background(), username)
	if err != nil {
		return false
	}
	return true
}

//sync to redis
func (ru roomUsecase) GetChatHistory(ctx context.Context, username1, username2, fromTS, toTS string) *presenter.ChatResponse {
	res := &presenter.ChatResponse{}

	fmt.Println(username1, username2)
	//check if user exist
	if !ru.VerifyContact(username1) || !re.VerifyContact(username2) {
		res.Message = "User not found"
		return res
	}

	chats, err := ru.roomRepo.FetchChatBetween(username1, username2, fromTS, toTS)
	if err != nil {
		log.Println("error in fetching chat history betweeen", err)
		res.Message = "unable to fetch chat history. please try again later"
		return res
	}

	res.Status = true
	res.Data = chats
	res.Total = len(chats)
	return res
}