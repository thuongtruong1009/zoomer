package usecase

import (
	"bytes"
	"github.com/minio/minio-go/v7"
	"zoomer/internal/models"
	"zoomer/internal/resources/adapter"
	"zoomer/internal/resources/repository"
)

func GetImage(Client *minio.Client, bucketName, objectName string) (res models.Resource) {
	todo := adapter.GetData(Client, bucketName, objectName)
	res = repository.GetResource(todo)
	return res
}

func GetAllImages(Client *minio.Client, bucketName string) (res models.ResourceList) {
	todoList := adapter.GetDataList(Client, bucketName)
	res = repository.GetResourcesList(todoList)
	return res
}

func AddImage(Client *minio.Client, bucketName, objectName, id, name string) (res models.Resource) {
	jsonData, res := repository.CreateResource(id, name)
	data := bytes.NewReader(jsonData)
	err := adapter.UploadData(Client, bucketName, objectName, data)
	if err != nil {
		panic(err)
	}
	return res
}

func UploadImage(Client *minio.Client, bucketName, objectName, id, name string) {
	jsonFile, _ := repository.CreateResource(id, name)
	data := bytes.NewReader(jsonFile)
	err := adapter.UploadData(Client, bucketName, objectName, data)
	if err != nil {
		panic(err)
	}
}

func DeleteImage(Client *minio.Client, bucketName, objectName string) {
	err := adapter.DeleteData(Client, bucketName, objectName)
	if err != nil {
		panic(err)
	}
}
