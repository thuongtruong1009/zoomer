package usecase

import (
	"context"
	"zoomer/internal/models"
	"zoomer/internal/auth/presenter"
)

type UseCase interface {
	SignUp(ctx context.Context, username, password string, limit int) (*models.User, error)

	SignIn(ctx context.Context, username, password string) (*presenter.LogInResponse, error)

	ParseToken(ctx context.Context, accessToken string) (string, error)
}
