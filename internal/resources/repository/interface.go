package repository

import (
	"io"
	"zoomer/internal/models"
)

type ResourceRepository interface {
	CreateResource(id, name string) (jsonData []byte, todo models.Resource)

	GetResource(jsonFile io.Reader) (temp models.Resource)

	GetResourcesList(jsonFiles []io.Reader) (temp models.ResourceList)
}
