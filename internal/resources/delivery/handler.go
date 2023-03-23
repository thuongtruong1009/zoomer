package delivery

import (
	"log"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"zoomer/internal/models"
	"zoomer/internal/resources/usecase"
)

type resourceHandler struct {
	resourceUC usecase.ResourceUseCase
}

func NewResourceHandler(resourceUC usecase.ResourceUseCase) *resourceHandler {
	return &resourceHandler{resourceUC: resourceUC}
}

func (rh *resourceHandler) GetResource(Client *minio.Client, bucketName string) echo.HandlerFunc{
	return func(c echo.Context) error {
		ab := rh.resourceUC.GetAllImages(Client, bucketName)
		return c.JSON(200, ab)
	}
}

func (rh *resourceHandler) CreateResource(Client *minio.Client, bucketName string) echo.HandlerFunc{
	return func(c echo.Context) error {
		var todo models.Resource
		err := c.Bind(&todo)
		if err != nil {
			log.Fatal(err)
		}
		res := rh.resourceUC.AddImage(Client, bucketName, c.Param("id")+".json", c.Param("id"), todo.Name)
		return c.JSON(200, res)
	}
}

//example data
	// UploadResource(Client, "todolist", "todo1.json", "1", "go to school")
	// UploadResource(Client, "todolist", "todo2.json", "2", "go to canteen")
	// UploadResource(Client, "todolist", "todo3.json", "3", "come back home")

func (rh *resourceHandler) UploadResource(Client *minio.Client, bucketName string) echo.HandlerFunc{
	return func(c echo.Context) error {
		var todo models.Resource
		err := c.Bind(&todo)
		if err != nil {
			log.Fatal(err)
		}
		rh.resourceUC.UploadImage(Client, bucketName, c.Param("id")+".json", c.Param("id"), todo.Name)
		return c.JSON(200, todo)
	}
}

func (rh *resourceHandler) DeleteResource(Client *minio.Client, bucketName string) echo.HandlerFunc{
	return func(c echo.Context) error {
		rh.resourceUC.DeleteImage(Client, bucketName, c.Param("id")+".json")
		return c.JSON(200, "deleted")
	}
}
