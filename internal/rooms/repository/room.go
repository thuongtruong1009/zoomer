package repository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"github.com/go-redis/redis/v8"
	"github.com/thuongtruong1009/zoomer/infrastructure/cache"
	"github.com/thuongtruong1009/zoomer/internal/models"
	chatAdapter "github.com/thuongtruong1009/zoomer/internal/chats/adapter"
	// "github.com/RediSearch/redisearch-go/redisearch"
)

type roomRepository struct {
	pgDB    *gorm.DB
	redisDB *redis.Client
}

func NewRoomRepository(pgDB *gorm.DB, redisDB *redis.Client) RoomRepository {
	return &roomRepository{
		pgDB:    pgDB,
		redisDB: redisDB,
	}
}

func (rr *roomRepository) CreateRoom(ctx context.Context, room *models.Room) error {
	result := rr.pgDB.WithContext(ctx).Create(&room)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (rr *roomRepository) GetRoomsByUserId(ctx context.Context, userId string) ([]*models.Room, error) {
	//check in cache
	UsernameInCache := cache.GetCache(cache.UserRoomKey(userId))
	if UsernameInCache != nil {
		return UsernameInCache.([]*models.Room), nil
	}

	var rooms []*models.Room
	err := rr.pgDB.WithContext(ctx).Where(&models.Room{
		CreatedBy: userId}).Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	//set to cache
	cache.SetCache(cache.UserRoomKey(userId), rooms, 0)

	return rooms, nil
}

func (rr *roomRepository) GetAllRooms(ctx context.Context) ([]*models.Room, error) {
	var chats []*models.Room
	err := rr.pgDB.WithContext(ctx).Limit(200).Find(&chats).Error

	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (rr *roomRepository) CountRooms(ctx context.Context, userId string) (int, error) {
	var count int

	err := rr.pgDB.WithContext(ctx).Raw(`SELECT COUNT(*) FROM rooms WHERE rooms.created_by = ? AND DATE_TRUNC('day', "created_at") = CURRENT_DATE GROUP BY DATE_TRUNC('day', "created_at")`, userId).Scan(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

// sync to redis
func (rr *roomRepository) FetchChatBetween(ctx context.Context, username1, username2, fromTS, toTS string) ([]models.Chat, error) {
	query := fmt.Sprintf("@from:{%s|%s} @to:{%s|%s} @timestamp:[%s %s]", username2, username1, username1, username2, fromTS, toTS)

	res, err := rr.redisDB.Do(ctx, "FT.SEARCH", chatAdapter.ChatIndex(), query, "SORTBY", "timestamp", "DESC").Result()
	if err != nil {
		return nil, err
	}

	fmt.Println("\nstep 1: \n", res)

	data := chatAdapter.Deserialise(res)

	fmt.Println("\nstep 2: \n", data)

	chats := chatAdapter.DeserialiseChat(data)

	return chats, nil
}

func (rr *roomRepository) FetchContactList(ctx context.Context, username string) ([]models.ContactList, error) {
	zRangeArg := redis.ZRangeArgs{
		Key:   chatAdapter.ContactListZKey(username),
		Start: 0,
		Stop:  -1,
		Rev:   true,
	}

	res, err := rr.redisDB.ZRangeArgsWithScores(ctx, zRangeArg).Result()
	if err != nil {
		return nil, err
	}
	contactList := chatAdapter.DeserialiseContactList(res)
	return contactList, nil
}

func (rr *roomRepository) CreateFetchChatBetweenIndex() {
	res, err := rr.redisDB.Do(context.Background(), "FT.CREATE", chatAdapter.ChatIndex(), "ON", "JSON",
		"PREFIX", "1", "chat#",
		"SCHEMA", "$.from", "AS", "from", "TAG",
		"$.to", "AS", "to", "TAG",
		"$.timestamp", "AS", "timestamp", "NUMERIC", "SORTABLE").Result()

	fmt.Println(res, err)
}
