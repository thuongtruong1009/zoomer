package delivery

import (
	"time"
	"log"
	"github.com/labstack/echo/v4"
	"zoomer/internal/resources/adapter"
)

func MapResourceRoutes(resourceGroup string, port string) {
	time.Sleep(3 * time.Second)
	Client, err := adapter.MinioClient()
	if err != nil {
		log.Println(err)
	}
	err = adapter.CreateBucket(Client, "todolist")
	if err != nil {
		log.Println(err)
	}
	
	bucketName := "todolist"
	router := echo.New()
	router.GET("/image", GetResource(Client, bucketName))
	router.POST("/image/:uid/:id", CreateResource(Client, bucketName))
	router.PUT("/image/:uid/:id", UploadResource(Client, bucketName))
	router.DELETE("/image/:uid/:id", DeleteResource(Client, bucketName))
	router.Logger.Fatal(router.Start(port))
}
