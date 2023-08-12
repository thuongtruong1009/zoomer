package repository

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
)

const CtxUserKey = "userId"

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error

	GetUserByUsername(ctx context.Context, username string) (*models.User, error)

	GetUserById(ctx context.Context, userId string) (*models.User, error)
}
