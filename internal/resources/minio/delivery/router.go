package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/internal/resources/minio/adapter"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"log"
	"time"
)

func MapResourceRoutes(resourceGroup *echo.Group, rh ResourceHandler) {
	time.Sleep(3 * time.Second)
	Client, err := adapter.MinioClient()
	if err != nil {
		log.Println(err)
	}

	err = adapter.CreateBucket(Client, constants.BucketName)
	if err != nil {
		log.Println(err)
	}

	resourceGroup.GET("/image", rh.GetResource(Client, constants.BucketName))
	resourceGroup.POST("/image/:uid/:id", rh.CreateResource(Client, constants.BucketName))
	resourceGroup.PUT("/image/:uid/:id", rh.UploadResource(Client, constants.BucketName))
	resourceGroup.DELETE("/image/:uid/:id", rh.DeleteResource(Client, constants.BucketName))
	// resourceGroup.Logger.Fatal(router.Start(port))
}
