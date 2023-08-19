package usecase

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/modules/users/presenter"
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

func (u *userUseCase) GetUserByIdOrName(ctx context.Context, IdOrName string) (*presenter.GetUserByIdOrNameResponse, error) {
	user, err := u.repo.GetUserByIdOrName(ctx, IdOrName)
	if err != nil {
		return nil, err
	}

	res := &presenter.GetUserByIdOrNameResponse{
		Id:       user.Id,
		Username: user.Username,
		Email: user.Email,
		Limit:    user.Limit,
	}
	return res, nil
}
