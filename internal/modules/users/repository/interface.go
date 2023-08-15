package repository

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
)

type IUserRepository interface{
	GetUserByIdOrName(ctx context.Context, IdOrUserName string) (*models.User, error)
}
