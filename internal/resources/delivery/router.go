package delivery

import (
	"time"
	"log"
	"github.com/labstack/echo/v4"
	"zoomer/internal/resources/adapter"
)

func MapResourceRoutes(resourceGroup *echo.Group, rh ResourceHandler) {
	time.Sleep(3 * time.Second)
	Client, err := adapter.MinioClient()
	if err != nil {
		log.Println(err)
	}
	bucketName := "todolist"

	err = adapter.CreateBucket(Client, bucketName)
	if err != nil {
		log.Println(err)
	}
	
	resourceGroup.GET("/image", rh.GetResource(Client, bucketName))
	resourceGroup.POST("/image/:uid/:id", rh.CreateResource(Client, bucketName))
	resourceGroup.PUT("/image/:uid/:id", rh.UploadResource(Client, bucketName))
	resourceGroup.DELETE("/image/:uid/:id", rh.DeleteResource(Client, bucketName))
	// resourceGroup.Logger.Fatal(router.Start(port))
}
