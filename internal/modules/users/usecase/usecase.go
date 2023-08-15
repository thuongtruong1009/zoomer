package usecase

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/modules/users/repository"
)

type userUseCase struct {
	repo repository.IUserRepository
}

func NewUserUseCase(repo repository.IUserRepository) IUserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (u *userUseCase) GetUserByIdOrName(ctx context.Context, IdOrName string) (*models.User, error) {
	return u.repo.GetUserByIdOrName(ctx, IdOrName)
}
