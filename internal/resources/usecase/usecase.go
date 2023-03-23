package resources

import (
	"bytes"
	"github.com/minio/minio-go/v7"
)

func GetAllTodoss(Client *minio.Client, bucketName string) (res TodoList) {
	todoList := GetDataList(Client, bucketName)
	res = GetAllTodos(todoList)
	return res
}

func AddTodo(Client *minio.Client, bucketName, objectName, id, name string) (res Todo) {
	jsonData, res := CreateJson(id, name)
	data := bytes.NewReader(jsonData)
	err := UploadData(Client, bucketName, objectName, data)
	if err != nil {
		panic(err)
	}
	return res
}

func UploadJson(Client *minio.Client, bucketName, objectName, id, name string) {
	jsonFile, _ := CreateJson(id, name)
	data := bytes.NewReader(jsonFile)
	err := UploadData(Client, bucketName, objectName, data)
	if err != nil {
		panic(err)
	}
}
