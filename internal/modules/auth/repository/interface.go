package repository

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/presenter"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error

	UpdatePassword(ctx context.Context, dto *presenter.ResetPassword) error
}
