package usecase

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/auth/presenter"
	"github.com/thuongtruong1009/zoomer/internal/models"
)

type UseCase interface {
	SignUp(ctx context.Context, username, password string, limit int) (*models.User, error)

	SignIn(ctx context.Context, username, password string) (*presenter.LogInResponse, error)

	ParseToken(ctx context.Context, accessToken string) (string, error)
}
