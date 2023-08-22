package usecase

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/presenter"
)

type UseCase interface {
	SignUp(ctx context.Context, dto *presenter.SignUpRequest) (*presenter.SignUpResponse, error)

	SignIn(ctx context.Context, dto *presenter.SignInRequest) (*presenter.SignInResponse, error)

	ParseToken(ctx context.Context, accessToken string) (string, error)

	ForgotPassword(ctx context.Context, dto *presenter.ForgotPassword) error

	ResetPassword(ctx context.Context, dto *presenter.ResetPassword) error
}
