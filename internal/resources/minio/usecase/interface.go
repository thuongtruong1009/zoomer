package usecase

import "github.com/thuongtruong1009/zoomer/internal/models"

type ResourceUseCase interface {
	GetImage(objectName string) models.Resource

	GetAllImages() models.ResourceList

	AddImage(objectName, id, name string) models.Resource

	UploadImage(objectName, id, name string)

	DeleteImage(objectName string)
}
