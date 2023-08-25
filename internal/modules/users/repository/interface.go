package repository

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/pkg/abstract"
)

type IUserRepository interface {
	GetUserByIdOrName(ctx context.Context, IdOrUserName string) (*models.User, error)
	Search(ctx context.Context, name string, pagination *abstract.Pagination) (*models.UsersList, error)
}
