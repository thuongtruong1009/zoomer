package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/thuongtruong1009/zoomer/infrastructure/cache"
	// chatAdapter "github.com/thuongtruong1009/zoomer/internal/modules/chats/adapter"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"gorm.io/gorm"
	// "log"
	"strings"
	"time"
	"fmt"
	"github.com/google/uuid"
)

type userRepository struct {
	pgDB *gorm.DB
	redisDB *redis.Client
	paramCfg *parameter.ParameterConfig
}

func NewUserRepository(pgDB *gorm.DB, redisDB *redis.Client, paramCfg *parameter.ParameterConfig) IUserRepository {
	return &userRepository {
		pgDB: pgDB,
		redisDB: redisDB,
		paramCfg: paramCfg,
	}
}

func (ur *userRepository) GetUserByIdOrName(ctx context.Context, IdOrUsername string) (*models.User, error) {
	var queryStruct *models.User
	var queryCacheKey string

	_, err := uuid.Parse(IdOrUsername)
	if err != nil {
		queryStruct = &models.User{
			Username: strings.ToLower(IdOrUsername),
		}
		queryCacheKey = cache.UsernameKey(IdOrUsername)
	} else {
		queryCacheKey = cache.UserIdKey(IdOrUsername)
		queryStruct = &models.User{
			Id: IdOrUsername,
		}
	}

	fmt.Println(queryStruct)
	fmt.Println(queryCacheKey)

	//check in cache
	userInCache := cache.GetCache(queryCacheKey)
	if userInCache != nil {
		return userInCache.(*models.User), nil
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, ur.paramCfg.OtherConf.CtxTimeout*time.Second)
	defer cancel()

	var user models.User
	if err := ur.pgDB.WithContext(timeoutCtx).Where(queryStruct).First(&user).Error; err != nil {
		return nil, constants.ErrNoRecord
	}

	//redis sync
	// _, err := ur.redisDB.SIsMember(context.Background(), chatAdapter.UserSetKey(), IdOrUsername).Result()
	// if err != nil {
	// 	log.Fatalln("(Redis) while checking user existance: ", err)
	// }
	// if !exist {
	// 	log.Println("(Redis) user not found in Redis-DB")
	// 	return nil, constants.ErrNoRecord
	// }

	//set in cache
	cache.SetCache(queryCacheKey, &user, 0)
	return &user, nil
}
