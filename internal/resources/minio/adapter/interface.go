package adapter

import "io"

type ResourceAdapter interface {
	GetData(objectName string) (file io.Reader)

	GetDataList() (file []io.Reader)

	UploadData(objectName string, file io.Reader) error

	DeleteData(objectName string) error

	DeleteDataList() error
}
