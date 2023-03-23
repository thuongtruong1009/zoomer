package usecase

import (
	"context"
	"io"
	"zoomer/internal/models"
)

type ResourceUseCase interface {
	UploadResource(ctx context.Context, file io.Reader, fileName string) (*models.Resource, error)

	GetResource(ctx context.Context, fileName string) (io.Reader, error)

	DeleteResource(ctx context.Context, fileName string) error
}
