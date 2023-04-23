package repository

import (
	"context"
	"gorm.io/gorm"
	"strings"
	"time"
	"zoomer/internal/models"
	"zoomer/pkg/constants"
	// "zoomer/pkg/cache"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if err := ur.db.WithContext(timeoutCtx).Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	//check in cache
	// UsernameInCache := cache.GetCache(cache.UsernameKey(username))
	// if UsernameInCache != nil {
	// 	return UsernameInCache.(*models.User), nil
	// }

	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var user models.User
	if err := ur.db.WithContext(timeoutCtx).Where(&models.User{
		Username: strings.ToLower(username),
	}).First(&user).Error; err != nil {
		return nil, constants.ErrNoRecord
	}

	//set in cache
	// cache.SetCache(cache.UsernameKey(username), &user, 0)

	return &user, nil
}

func (ur *userRepository) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	//check in cache
	// UsernameInCache := cache.GetCache(cache.UserIdKey(userId))
	// if UsernameInCache != nil {
	// 	return UsernameInCache.(*models.User), nil
	// }

	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var user models.User
	if err := ur.db.WithContext(timeoutCtx).Where(&models.User{
		Id: userId,
	}).First(&user).Error; err != nil {
		return nil, constants.ErrNoRecord
	}

	// set in cache
	// cache.SetCache(cache.UserIdKey(userId), &user, 0)

	return &user, nil
}

func (ur *userRepository) QueryMatchingFields(ctx context.Context, match string) (*[]models.User, error) {
	var user []models.User
	err := ur.db.WithContext(ctx).Where("username LIKE ?", "%"+match+"%").First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
