package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

type ResourceHandler interface {
	GetResource(Client *minio.Client, bucketName string) echo.HandlerFunc

	CreateResource(Client *minio.Client, bucketName string) echo.HandlerFunc

	UploadResource(Client *minio.Client, bucketName string) echo.HandlerFunc

	DeleteResource(Client *minio.Client, bucketName string) echo.HandlerFunc
}