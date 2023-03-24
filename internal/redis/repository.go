package repository

import (
	"example.com/myapp/domain"
	"example.com/myapp/redis"
)

type redisRepository struct {
	db    *sql.DB
	cache redis.Cache
}

func NewRedisRepository(db *sql.DB, cache redis.Cache) domain.Repository {
	return &redisRepository{
		db:    db,
		cache: cache,
	}
}

func (r *redisRepository) GetByID(id int) (*domain.Entity, error) {
	key := fmt.Sprintf("entity:%d", id)
	cached, err := r.cache.Get(key)
	if err == nil {
		entity := &domain.Entity{}
		err := json.Unmarshal([]byte(cached), entity)
		if err == nil {
			return entity, nil
		}
	}

	bytes, err := json.Marshal(entity)
	if err == nil {
		r.cache.Set(key, string(bytes), 24*time.Hour)
	}

	return entity, nil
}
