package usecase

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/models"
)

type IUserUseCase interface {
	GetUserByIdOrName(ctx context.Context, IdOrName string) (*models.User, error)
}
