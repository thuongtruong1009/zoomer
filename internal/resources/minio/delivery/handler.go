package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/internal/resources/minio/usecase"
	"log"
)

type resourceHandler struct {
	resourceUC usecase.ResourceUseCase
}

func NewResourceHandler(resourceUC usecase.ResourceUseCase) *resourceHandler {
	return &resourceHandler{resourceUC: resourceUC}
}

func (rh *resourceHandler) GetResource() echo.HandlerFunc {
	return func(c echo.Context) error {
		ab := rh.resourceUC.GetAllImages()
		return c.JSON(200, ab)
	}
}

func (rh *resourceHandler) CreateResource() echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo models.Resource
		err := c.Bind(&todo)
		if err != nil {
			log.Fatal(err)
		}
		res := rh.resourceUC.AddImage(c.Param("uid"), c.Param("id"), todo.Name)
		return c.JSON(200, res)
	}
}

//example data
// UploadResource(Client, "todolist", "todo1.json", "1", "go to school")
// UploadResource(Client, "todolist", "todo2.json", "2", "go to canteen")
// UploadResource(Client, "todolist", "todo3.json", "3", "come back home")

func (rh *resourceHandler) UploadResource() echo.HandlerFunc {
	return func(c echo.Context) error {
		var todo models.Resource
		err := c.Bind(&todo)
		if err != nil {
			log.Fatal(err)
		}
		rh.resourceUC.UploadImage(c.Param("id")+".json", c.Param("id"), todo.Name)
		return c.JSON(200, todo)
	}
}

func (rh *resourceHandler) DeleteResource() echo.HandlerFunc {
	return func(c echo.Context) error {
		rh.resourceUC.DeleteImage(c.Param("id") + ".json")
		return c.JSON(200, "deleted")
	}
}
