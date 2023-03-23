package adapter

import (
	"io"
	"github.com/minio/minio-go/v7"
)

type ResourceAdapter interface {
	GetData(client *minio.Client, bucketName, objectName string) (file io.Reader)

	GetDataList(client *minio.Client, bucketName string) (file []io.Reader)

	UploadData(client *minio.Client, bucketName, objectName string, file io.Reader) error

	DeleteData(client *minio.Client, bucketName, objectName string) error

	DeleteDataList(client *minio.Client, bucketName string) error
}