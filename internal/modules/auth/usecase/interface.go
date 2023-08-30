package usecase

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/modules/auth/presenter"
)

type UseCase interface {
	ParseToken(ctx context.Context, accessToken string) (*models.AuthClaims, error)

	SignUp(ctx context.Context, dto *presenter.SignUpRequest) (*presenter.SignUpResponse, error)

	SignIn(ctx context.Context, dto *presenter.SignInRequest) (*presenter.SignInResponse, error)

	ForgotPassword(ctx context.Context, email string) error

	VerifyResetPasswordOtp(ctx context.Context, otpCode string) error

	ResetPassword(ctx context.Context, dto *presenter.ResetPassword) error
}
