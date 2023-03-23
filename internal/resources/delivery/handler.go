package delivery

import (
	"log"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"zoomer/internal/models"
	"zoomer/internal/resources/usecase"
)

func GetResource(Client *minio.Client, bucketName string) echo.HandlerFunc{
	return func(c echo.Context) error {
		ab := usecase.GetAllImages(Client, bucketName)
		return c.JSON(200, ab)
	}
}

func CreateResource(Client *minio.Client, bucketName string) echo.HandlerFunc{
	return func(c echo.Context) error {
		var todo models.Resource
		err := c.Bind(&todo)
		if err != nil {
			log.Fatal(err)
		}
		res := usecase.AddImage(Client, bucketName, c.Param("id")+".json", c.Param("id"), todo.Name)
		return c.JSON(200, res)
	}
}

//example data
	// UploadResource(Client, "todolist", "todo1.json", "1", "go to school")
	// UploadResource(Client, "todolist", "todo2.json", "2", "go to canteen")
	// UploadResource(Client, "todolist", "todo3.json", "3", "come back home")

func UploadResource(Client *minio.Client, bucketName string) echo.HandlerFunc{
	return func(c echo.Context) error {
		var todo models.Resource
		err := c.Bind(&todo)
		if err != nil {
			log.Fatal(err)
		}
		usecase.UploadImage(Client, bucketName, c.Param("id")+".json", c.Param("id"), todo.Name)
		return c.JSON(200, todo)
	}
}

func DeleteResource(Client *minio.Client, bucketName string) echo.HandlerFunc{
	return func(c echo.Context) error {
		usecase.DeleteImage(Client, bucketName, c.Param("id")+".json")
		return c.JSON(200, "deleted")
	}
}
