package usecase

import (
	"bytes"
	"fmt"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/modules/resources/minio/adapter"
	"github.com/thuongtruong1009/zoomer/internal/modules/resources/minio/repository"
)

type resourceUsecase struct {
	resourceAdapter adapter.ResourceAdapter
	resourceRepo    repository.ResourceRepository
}

func NewMinioResourceUseCase(resourceAdapter adapter.ResourceAdapter, resourceRepo repository.ResourceRepository) ResourceUseCase {
	return &resourceUsecase{
		resourceAdapter: resourceAdapter,
		resourceRepo:    resourceRepo,
	}
}

func (ru resourceUsecase) GetImage(objectName string) models.Resource {
	todo := ru.resourceAdapter.GetData(objectName)
	res := ru.resourceRepo.GetResource(todo)

	return res
}

func (ru resourceUsecase) GetAllImages() models.ResourceList {
	todoList := ru.resourceAdapter.GetDataList()
	res := ru.resourceRepo.GetResourcesList(todoList)
	fmt.Println("res 1", todoList)
	fmt.Println("res 2", res)
	return res
}

func (ru resourceUsecase) AddImage(objectName, id, name string) models.Resource {
	jsonData, res := ru.resourceRepo.CreateResource(id, name)
	data := bytes.NewReader(jsonData)
	err := ru.resourceAdapter.UploadData(objectName, data)
	if err != nil {
		panic(err)
	}
	return res
}

func (ru resourceUsecase) UploadImage(objectName, id, name string) {
	jsonFile, _ := ru.resourceRepo.CreateResource(id, name)
	data := bytes.NewReader(jsonFile)
	err := ru.resourceAdapter.UploadData(objectName, data)
	if err != nil {
		panic(err)
	}
}

func (ru resourceUsecase) DeleteImage(objectName string) {
	err := ru.resourceAdapter.DeleteData(objectName)
	if err != nil {
		panic(err)
	}
}
