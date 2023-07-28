package repository

import (
	"context"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
	"github.com/go-redis/redis/v8"
	chatAdapter "github.com/thuongtruong1009/zoomer/internal/chats/adapter"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/pkg/cache"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

type authRepository struct {
	pgDB    *gorm.DB
	redisDB *redis.Client
}

func NewAuthRepository(pgDB *gorm.DB, redisDB *redis.Client) UserRepository {
	return &authRepository{
		pgDB:    pgDB,
		redisDB: redisDB,
	}
}

func (ar *authRepository) CreateUser(ctx context.Context, user *models.User) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if err := ar.pgDB.WithContext(timeoutCtx).Create(&user).Error; err != nil {
		return err
	}

	//redis sync
	err := ar.redisDB.Set(context.Background(), user.Username, user.Password, 0).Err()
	if err != nil {
		log.Println("(Redis) while adding new user: ", err)
		return err
	}

	err = ar.redisDB.SAdd(context.Background(), chatAdapter.UserSetKey(), user.Username).Err()
	if err != nil {
		log.Println("(Redis) while adding new user: ", err)
		ar.redisDB.Del(context.Background(), user.Username)
		return err
	}

	return nil
}

func (ar *authRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	// check in cache
	UsernameInCache := cache.GetCache(cache.UsernameKey(username))
	if UsernameInCache != nil {
		return UsernameInCache.(*models.User), nil
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var user models.User
	if err := ar.pgDB.WithContext(timeoutCtx).Where(&models.User{
		Username: strings.ToLower(username),
	}).First(&user).Error; err != nil {
		return nil, constants.ErrNoRecord
	}

	//redis sync
	exist, err := ar.redisDB.SIsMember(context.Background(), chatAdapter.UserSetKey(), username).Result()
	if err != nil {
		log.Fatalln("(Redis) while checking user existance: ", err)
	}

	if !exist {
		log.Println("(Redis) user not found in Redis-DB")
		return nil, constants.ErrNoRecord
	}

	//set in cache
	cache.SetCache(cache.UsernameKey(username), &user, 0)

	return &user, nil
}

func (ar *authRepository) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	//check in cache
	UsernameInCache := cache.GetCache(cache.UserIdKey(userId))
	if UsernameInCache != nil {
		return UsernameInCache.(*models.User), nil
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var user models.User
	if err := ar.pgDB.WithContext(timeoutCtx).Where(&models.User{
		Id: userId,
	}).First(&user).Error; err != nil {
		return nil, constants.ErrNoRecord
	}

	// set in cache
	cache.SetCache(cache.UserIdKey(userId), &user, 0)

	return &user, nil
}

func (ar *authRepository) QueryMatchingFields(ctx context.Context, match string) (*[]models.User, error) {
	var user []models.User
	err := ar.pgDB.WithContext(ctx).Where("username LIKE ?", "%"+match+"%").First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
