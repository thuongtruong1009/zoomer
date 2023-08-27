package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/thuongtruong1009/zoomer/infrastructure/cache"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"gorm.io/gorm"
	"strings"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"github.com/thuongtruong1009/zoomer/pkg/abstract"
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

	_, err := uuid.Parse(account)
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

	// Count the total number of matching users
	if err := ur.pgDB.WithContext(ctx).Model(&models.User{}).Where("username LIKE ?", "%"+name+"%").Count(&total).Error; err != nil {
		return nil, err // Return the actual error instead of constants.ErrNoRecord
	}

	// Retrieve paginated users list
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



	// searchWord := fmt.Sprint("%s:*", name)
	// var count int
	// if err := ur.pgDB.QueryRow(ctx, `SELECT count(email_id) FROM users WHERE document_with_idx @@ to_tsquery($1)`, searchWord).Scan(&count); err != nil {
	// 	return nil, err
	// }

	// if count == 0 {
	// 	return &models.UsersList {
	// 		TotalCount: 0,
	// 		TotalPages: 0,
	// 		Page: 0,
	// 		Size: 0,
	// 		HasMore: false,
	// 		Users: make([]*models.Email, 0),
	// 	}, nil
	// }

	// rows, err := ur.pgDB.Query(ctx, `SELECT email_id, address_to, address_from, subject, message, created_at
	// FROM emails WHERE document_with_idx @@ to_tsquery($1) ORDER BY created_at OFFSET $2 LIMIT $3`, searchWord, pagination.GetOffset(), Pagination.GetLimit())
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()

	// userList := make([]*models.Email, 0, count)
	// for rows.Next() {
	// 	var m models.User
	// 	if err := rows.Scan(&m.Id, &m.Email, &m.Username); err != nil {
	// 		return nil, err
	// 	}
	// 	userList = append(userList, &m)
	// }

	// if err := rows.Err(); err != nil {
	// 	return nil, err
	// }

	// return &models.UsersList {
	// 	TotalCount: int64(count),
	// 	TotalPages: int64(pagination.GetTotalPages(count)),
	// 	Page: int64(pagination.GetPage()),
	// 	Size: int64(pagination.GetSize()),
	// 	HasMore: pagination.GetHasMore(count),
	// 	Users: userList,
	// }, nil
}
