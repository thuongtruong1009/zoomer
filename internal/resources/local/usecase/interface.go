package usecase

import (
	"context"
	"mime/multipart"
)

type ILocalResourceUseCase interface {
	UploadSingleFile(ctx context.Context, file *multipart.FileHeader) (string, error)
	UploadMultipleFile(ctx context.Context, files []*multipart.FileHeader) ([]string, error)
	DeleteSingleFile(ctx context.Context, fileName string) error
	DeleteMultipleFile(ctx context.Context, fileNames []string) error
}


