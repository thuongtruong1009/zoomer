package repository

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error

	UpdatePassword(ctx context.Context, email, newPassword string) error
}
