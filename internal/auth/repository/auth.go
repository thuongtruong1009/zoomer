package repository

import (
	"context"
	"gorm.io/gorm"
	"strings"
	"zoomer/internal/models"
	// "zoomer/pkg/cache"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	result := ur.db.WithContext(ctx).Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	//check in cache
	// UsernameInCache := cache.GetCache(cache.UsernameKey(username))
	// if UsernameInCache != nil {
	// 	return UsernameInCache.(*models.User), nil
	// }

	var user models.User
	err := ur.db.WithContext(ctx).Where(&models.User{
		Username: strings.ToLower(username),
	}).First(&user).Error

	if err != nil {
		return nil, err
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

	var user models.User
	err := ur.db.WithContext(ctx).Where(&models.User{
		Id: userId,
	}).First(&user).Error

	if err != nil {
		return nil, err
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
