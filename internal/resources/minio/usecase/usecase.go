package usecase

import (
	"bytes"
	"github.com/minio/minio-go/v7"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/resources/minio/adapter"
	"github.com/thuongtruong1009/zoomer/internal/resources/minio/repository"
)

type resourceUsecase struct {
	resourceRepo    repository.ResourceRepository
	resourceAdapter adapter.ResourceAdapter
}

func NewResourceUseCase(resourceRepo repository.ResourceRepository) ResourceUseCase {
	return &resourceUsecase{
		resourceRepo: resourceRepo,
	}
}

func (ru resourceUsecase) GetImage(Client *minio.Client, bucketName, objectName string) (res models.Resource) {
	todo := ru.resourceAdapter.GetData(Client, bucketName, objectName)
	res = ru.resourceRepo.GetResource(todo)
	return res
}

func (ru resourceUsecase) GetAllImages(Client *minio.Client, bucketName string) (res models.ResourceList) {
	todoList := ru.resourceAdapter.GetDataList(Client, bucketName)
	res = ru.resourceRepo.GetResourcesList(todoList)
	return res
}

func (ru resourceUsecase) AddImage(Client *minio.Client, bucketName, objectName, id, name string) (res models.Resource) {
	jsonData, res := ru.resourceRepo.CreateResource(id, name)
	data := bytes.NewReader(jsonData)
	err := ru.resourceAdapter.UploadData(Client, bucketName, objectName, data)
	if err != nil {
		panic(err)
	}
	return res
}

func (ru resourceUsecase) UploadImage(Client *minio.Client, bucketName, objectName, id, name string) {
	jsonFile, _ := ru.resourceRepo.CreateResource(id, name)
	data := bytes.NewReader(jsonFile)
	err := ru.resourceAdapter.UploadData(Client, bucketName, objectName, data)
	if err != nil {
		panic(err)
	}
}

func (ru resourceUsecase) DeleteImage(Client *minio.Client, bucketName, objectName string) {
	err := ru.resourceAdapter.DeleteData(Client, bucketName, objectName)
	if err != nil {
		panic(err)
	}
}
