package usecase

import (
	"context"
	"github.com/google/uuid"
	"io"
	"time"
	auth "zoomer/internal/auth/repository"
	"zoomer/internal/models"
	"zoomer/internal/resources/repository"
)

type resourceUseCase struct {
	resourceRepo repository.ResourceRepository
	userRepo     auth.UserRepository
}

func NewResourceUseCase(resourceRepo repository.ResourceRepository, userRepo auth.UserRepository) ResourceUseCase {
	return &resourceUseCase{
		resourceRepo: resourceRepo,
		userRepo:     userRepo,
	}
}

func (ru resourceUseCase) UploadResource(ctx context.Context, userId string, file io.Reader, fileName string) (*models.Resource, error) {
	resource := &models.Resource{
		Id:        uuid.New().String(),
		Name:      fileName,
		CreatedAt: time.Now(),
		CreatedBy: userId,
	}

	result, err := ru.resourceRepo.UploadResource(ctx, resource)
	if err != nil {
		return nil, err
	}
}

func (ru resourceUseCase) DownloadResource(ctx context.Context, fileName string) (io.Reader, error) {
	resource, err := ru.resourceRepo.DownloadResource(ctx, fileName)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (ru resourceUseCase) DeleteResource(ctx context.Context, fileName string) error {
	err := ru.resourceRepo.DeleteResource(ctx, fileName)
	if err != nil {
		return err
	}

	return nil
}

func (ru resourceUseCase) GetResourceById(ctx context.Context, resourceId string) (*models.Resource, error) {
	resource, err := ru.resourceRepo.GetResourceById(ctx, resourceId)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (re resourceUseCase) GetAllResourcesByBucketId(ctx context.Context, bucketId string) ([]*models.Resource, error) {
	resources, err := re.resourceRepo.GetAllResourcesByBucketId(ctx, bucketId)
	if err != nil {
		return nil, err
	}

	return resources, nil
}
