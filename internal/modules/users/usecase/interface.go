package usecase

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/modules/users/presenter"
)

type IUserUseCase interface {
	GetUserByIdOrName(ctx context.Context, IdOrName string) (*presenter.GetUserByIdOrNameResponse, error)
}
