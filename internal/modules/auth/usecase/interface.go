package usecase

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/presenter"
)

type UseCase interface {
	SignUp(ctx context.Context, username, password string, limit int) (*presenter.SignUpResponse, error)

	SignIn(ctx context.Context, username, password string) (*presenter.SignInResponse, error)

	ParseToken(ctx context.Context, accessToken string) (string, error)
}
