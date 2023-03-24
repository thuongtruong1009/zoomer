package usecase

import (
"example.com/myapp/redis"
)

type entityUsecase struct {
repo domain.Repository
cache redis.Cache
}

func NewEntityUsecase(repo domain.Repository, cache redis.Cache) domain.EntityUsecase {
return &entityUsecase{
repo: repo,
cache: cache,
}
}

func (u *entityUsecase) GetByID(id int) (*domain.Entity, error) {
entity, err := u.cache.Get(id)
if err == nil {
return entity, nil
}

entity, err = u.repo.GetByID(id)
if err != nil {
    return nil, err
}

u.cache.Set(id, entity, 24*time.Hour)

return entity, nil
