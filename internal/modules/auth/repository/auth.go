package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/internal/models"
	chatAdapter "github.com/thuongtruong1009/zoomer/internal/modules/chats/adapter"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"gorm.io/gorm"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

type authRepository struct {
	pgDB     *gorm.DB
	redisDB  *redis.Client
	paramCfg *parameter.ParameterConfig
}

func NewAuthRepository(pgDB *gorm.DB, redisDB *redis.Client, paramCfg *parameter.ParameterConfig) UserRepository {
	return &authRepository{
		pgDB:     pgDB,
		redisDB:  redisDB,
		paramCfg: paramCfg,
	}
}

func (ar *authRepository) CreateUser(ctx context.Context, user *models.User) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, helpers.DurationSecond(ar.paramCfg.CtxTimeout))
	defer cancel()

	if err := ar.pgDB.WithContext(timeoutCtx).Create(&user).Error; err != nil {
		exceptions.Log(constants.ErrorContextTimeout, err)
		return err
	}

	//redis sync
	err := ar.redisDB.Set(context.Background(), user.Username, user.Password, 0).Err()
	if err != nil {
		exceptions.Log(constants.ErrRedisSyncUser, err)
		return constants.ErrRedisSyncUser
	}

	err = ar.redisDB.SAdd(context.Background(), chatAdapter.UserSetKey(), user.Username).Err()
	if err != nil {
		exceptions.Log(constants.ErrRedisAddUser, err)
		ar.redisDB.Del(context.Background(), user.Username)
		return constants.ErrRedisAddUser
	}

	return nil
}

// func (ar *authRepository) QueryMatchingFields(ctx context.Context, match string) (*[]models.User, error) {
// 	var user []models.User
// 	err := ar.pgDB.WithContext(ctx).Where("username LIKE ?", "%"+match+"%").First(&user).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }
