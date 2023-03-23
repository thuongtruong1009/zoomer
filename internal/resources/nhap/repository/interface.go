package repository

import (
	"io"
)

type ResourceRepository interface {
	CreateBucket(bucketName string) error

	UploadResource(bucketName string, fileName string, file io.Reader) error

	GetResource(bucketName string, fileName string) (io.Reader, error)

	GetResourcesList(bucketName string) ([]string, error)

	DeleteResource(bucketName string, fileName string) error

	DeleteResourcesList(bucketName string) error
}
