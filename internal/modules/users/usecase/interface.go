package usecase

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/modules/users/presenter"
	"github.com/thuongtruong1009/zoomer/pkg/abstract"
)

type IUserUseCase interface {
	GetUserByIdOrName(ctx context.Context, IdOrName string) (*presenter.GetUserByIdOrNameResponse, error)
	SearchUser(ctx context.Context, name string, pagination *abstract.Pagination) (*models.UsersList, error)
}
