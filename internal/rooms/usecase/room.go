package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
	auth "github.com/thuongtruong1009/zoomer/internal/auth/repository"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/rooms/presenter"
	"github.com/thuongtruong1009/zoomer/internal/rooms/repository"
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
	_, err := ru.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return false
	}
	return true
}

// sync to redis
func (ru roomUsecase) GetChatHistory(ctx context.Context, username1, username2, fromTS, toTS string) *presenter.ChatResponse {
	res := &presenter.ChatResponse{}

	//check if user exist
	if !ru.VerifyContact(ctx, username1) || !ru.VerifyContact(ctx, username2) {
		res.Message = "(redis) User not found in Redis-DB"
		return res
	}

	chats, err := ru.roomRepo.FetchChatBetween(ctx, username1, username2, fromTS, toTS)
	if err != nil {
		fmt.Println("(redis) error in fetching chat history betweeen", err)
		res.Message = "(redis) unable to fetch chat history. please try again later"
		return res
	}


	res.Status = true
	res.Data = chats
	res.Total = len(chats)
	return res
}

func (ru roomUsecase) ContactList(ctx context.Context, username string) *presenter.ChatResponse {
	res := &presenter.ChatResponse{}

	if !ru.VerifyContact(ctx, username) {
		res.Message = "(redis) User not found"
		return res
	}

	contacts, err := ru.roomRepo.FetchContactList(ctx, username)
	if err != nil {
		fmt.Println("(redis) error in fetching contact list or username: ", err)
		res.Message = "(redis) unable to fetch contact list. please try again later"
		return res
	}

	res.Status = true
	res.Data = contacts
	res.Total = len(contacts)
	return res
}

func (ru roomUsecase) GetFetchChatBetweenIndex(ctx context.Context) {
	ru.roomRepo.CreateFetchChatBetweenIndex(ctx)
}
