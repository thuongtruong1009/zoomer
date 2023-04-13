package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"zoomer/db"
	chatAdapter "zoomer/internal/chats/adapter"
	"zoomer/internal/models"
	"zoomer/lib/cache"
)

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (cr *roomRepository) CreateRoom(ctx context.Context, room *models.Room) error {
	result := cr.db.WithContext(ctx).Create(&room)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *roomRepository) GetRoomsByUserId(ctx context.Context, userId string) ([]*models.Room, error) {
	//check in cache
	UsernameInCache := cache.GetCache(cache.UserRoomKey(userId))
	if UsernameInCache != nil {
		return UsernameInCache.([]*models.Room), nil
	}

	var rooms []*models.Room
	err := cr.db.WithContext(ctx).Where(&models.Room{
		CreatedBy: userId}).Find(&rooms).Error

	if err != nil {
		return nil, err
	}

	//set to cache
	cache.SetCache(cache.UserRoomKey(userId), rooms, 0)

	return rooms, nil
}

func (cr *roomRepository) GetAllRooms(ctx context.Context) ([]*models.Room, error) {
	var chats []*models.Room
	err := cr.db.WithContext(ctx).Limit(200).Find(&chats).Error

	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (cr *roomRepository) CountRooms(ctx context.Context, userId string) (int, error) {
	var count int

	err := cr.db.WithContext(ctx).Raw(`SELECT COUNT(*) FROM rooms WHERE rooms.created_by = ? AND DATE_TRUNC('day', "created_at") = CURRENT_DATE GROUP BY DATE_TRUNC('day', "created_at")`, userId).Scan(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

// sync to redis
func (cr *roomRepository) FetchChatBetween(ctx context.Context, username1, username2, fromTS, toTS string) ([]models.Chat, error) {
	query := fmt.Sprintf("@from:{%s|%s} @to:{%s|%s} @timestamp:[%s %s]", username1, username2, username1, username2, fromTS, toTS)

	res, err := db.GetRedisInstance().Do(ctx, "FT.SEARCH", chatAdapter.ChatIndex(), query, "SORTBY", "timestamp", "DESC").Result()

	if err != nil {
		return nil, err
	}

	data := chatAdapter.Deserialise(res)
	chats := chatAdapter.DeserialiseChat(data)
	return chats, nil
}

func (cr *roomRepository) FetchContactList(ctx context.Context, username string) ([]models.ContactList, error) {
	zRangeArg := redis.ZRangeArgs{
		Key:   chatAdapter.ContactListZKey(username),
		Start: 0,
		Stop:  -1,
		Rev:   true,
	}

	res, err := db.GetRedisInstance().ZRangeArgsWithScores(ctx, zRangeArg).Result()
	if err != nil {
		return nil, err
	}
	contactList := chatAdapter.DeserialiseContactList(res)
	return contactList, nil
}

func (cr *roomRepository) CreateFetchChatBetweenIndex(ctx context.Context) {
	res, err := db.GetRedisInstance().Do(ctx, "FT.CREATE", chatAdapter.ChatIndex(), "ON", "JSON", "PREFIX", "1", "chat#", "SCHEMA", "$.from", "AS", "from", "TAG", "$.to", "TAG", "$.timestamp", "AS", "timestamp", "NUMERIC", "SORTABLE").Result()

	fmt.Println(res, err)
}
