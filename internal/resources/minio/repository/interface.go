package repository

import (
	"github.com/thuongtruong1009/zoomer/internal/models"
	"io"
)

type ResourceRepository interface {
	CreateResource(id, name string) (jsonData []byte, todo models.Resource)

	GetResource(jsonFile io.Reader) (temp models.Resource)

	GetResourcesList(jsonFiles []io.Reader) (temp models.ResourceList)
}
