package repository

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
)

const CtxUserKey = "userId"

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error

	// CtxUserKey(ctx context.Context) string
}
