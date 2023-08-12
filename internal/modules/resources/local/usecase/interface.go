package usecase

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/modules/resources/local/presenter"
	"mime/multipart"
)

type ILocalResourceUseCase interface {
	UploadSingleFile(ctx context.Context, file *multipart.FileHeader) (*presenter.SingleUploadResponse, error)
	UploadMultipleFile(ctx context.Context, files []*multipart.FileHeader) (*presenter.MultipleUploadResponse, error)
	DeleteSingleFile(ctx context.Context, fileName string) error
	DeleteMultipleFile(ctx context.Context, fileNames []string) error
}
