package delivery

import (
	"github.com/labstack/echo/v4"
)

func MapResourceRoutes(resourceGroup string, port string) {
	bucketName := "todolist"
	router := echo.New()
	router.GET("/image", GetResource(bucketName))
	router.POST("/image/:uid/:id", CreateResource(bucketName))
	router.Logger.Fatal(router.Start(port))
}
