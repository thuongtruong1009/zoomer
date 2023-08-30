package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/thuongtruong1009/zoomer/infrastructure/cache"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/pkg/abstract"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"github.com/thuongtruong1009/zoomer/pkg/pipe"
	"gorm.io/gorm"
	"strings"
)

type userRepository struct {
	pgDB     *gorm.DB
	redisDB  *redis.Client
	paramCfg *parameter.ParameterConfig
}

func NewUserRepository(pgDB *gorm.DB, redisDB *redis.Client, paramCfg *parameter.ParameterConfig) IUserRepository {
	return &userRepository{
		pgDB:     pgDB,
		redisDB:  redisDB,
		paramCfg: paramCfg,
	}
}

func (ur *userRepository) GetUserByIdOrName(ctx context.Context, account string) (*models.User, error) {
	var queryStruct *models.User
	var queryCacheKey string

	err := pipe.IsValidUUID(account)

	if err != nil {
		isEmail := strings.IndexByte(account, '@')
		if isEmail >= 0 {
			queryStruct = &models.User{
				Email: account,
			}
		} else {
			queryStruct = &models.User{
				Username: strings.ToLower(account),
			}
		}
	} else {
		queryStruct = &models.User{
			Id: account,
		}
	}

	queryCacheKey = cache.AuthUserKey(account)

	userInCache := cache.GetCache(queryCacheKey)
	if userInCache != nil {
		return userInCache.(*models.User), nil
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, helpers.DurationSecond(ur.paramCfg.TokenTimeout))
	defer cancel()

	var user models.User
	if err := ur.pgDB.WithContext(timeoutCtx).Where(queryStruct).First(&user).Error; err != nil {
		exceptions.Log(constants.ErrorContextTimeout, err)
		return nil, constants.ErrNoRecord
	}

	cache.SetCache(queryCacheKey, &user, helpers.DurationSecond(ur.paramCfg.TokenTimeout))
	return &user, nil
}

func (ur *userRepository) Search(ctx context.Context, name string, pagination *abstract.Pagination) (*models.UsersList, error) {
	var users []*models.User
	var total int64

	if err := ur.pgDB.WithContext(ctx).Model(&models.User{}).Where("username LIKE ?", "%"+name+"%").Count(&total).Error; err != nil {
		return nil, err // Return the actual error instead of constants.ErrNoRecord
	}

	if err := ur.pgDB.WithContext(ctx).Where("username LIKE ?", "%"+name+"%").Limit(pagination.GetSize()).Offset(pagination.GetOffset()).Find(&users).Error; err != nil {
		return nil, err // Return the actual error instead of constants.ErrNoRecord
	}

	return &models.UsersList{
		TotalCount: total,
		TotalPages: int64(pagination.GetTotalPages(int(total))),
		Page:       int64(pagination.GetPage()),
		Size:       int64(pagination.GetSize()),
		HasMore:    pagination.GetHasMore(int(total)),
		Users:      users,
	}, nil
}
