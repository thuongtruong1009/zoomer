package usecase

import (
	"github.com/minio/minio-go/v7"
	"zoomer/internal/models"
)

type ResourceUseCase interface {
	GetImage(Client *minio.Client, bucketName, objectName string) (res models.Resource)

	GetAllImages(Client *minio.Client, bucketName string) (res models.ResourceList)

	AddImage(Client *minio.Client, bucketName, objectName, id, name string) (res models.Resource)

	UploadImage(Client *minio.Client, bucketName, objectName, id, name string)

	DeleteImage(Client *minio.Client, bucketName, objectName string)
}
