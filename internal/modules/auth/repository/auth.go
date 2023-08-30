package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/internal/models"
	chatAdapter "github.com/thuongtruong1009/zoomer/internal/modules/chats/adapter"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"gorm.io/gorm"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/presenter"
	userRepository "github.com/thuongtruong1009/zoomer/internal/modules/users/repository"
)

type authRepository struct {
	pgDB     *gorm.DB
	redisDB  *redis.Client
	paramCfg *parameter.ParameterConfig
	userRepo userRepository.IUserRepository
}

func NewAuthRepository(pgDB *gorm.DB, redisDB *redis.Client, paramCfg *parameter.ParameterConfig, userRepo userRepository.IUserRepository) UserRepository {
	return &authRepository{
		pgDB:     pgDB,
		redisDB:  redisDB,
		paramCfg: paramCfg,
		userRepo: userRepo,
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

func (ar *authRepository) UpdatePassword(ctx context.Context, dto *presenter.ResetPassword) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, helpers.DurationSecond(ar.paramCfg.CtxTimeout))
	defer cancel()

    user, err := ar.userRepo.GetUserByIdOrName(ctx, dto.Email)
	if err != nil {
		return err
	}

	hashedPassword, err := helpers.Encrypt(dto.NewPassword)
	if err != nil {
		exceptions.Log(constants.ErrHashPassword, err)
		return err
	}

	if (hashedPassword == user.Password) {
		exceptions.Log(constants.ErrPasswordNotChange, nil)
		return constants.ErrPasswordNotChange
	}

	if err := ar.pgDB.WithContext(timeoutCtx).Where("email = ?", dto.Email).Update("password", hashedPassword).Error; err != nil {
		exceptions.Log(constants.ErrorContextTimeout, err)
		return err
	}

	return nil
}
