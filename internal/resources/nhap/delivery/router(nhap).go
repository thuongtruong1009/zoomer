package delivery

import (
	"github.com/labstack/echo/v4"
)

func MapResourceRoutes(resourceGroup *echo.Group, h ResourceHandler) {
	resourceGroup.POST("/image", h.UploadImage())
	resourceGroup.GET("/image/:imageId", h.DownloadImage())
	resourceGroup.DELETE("/image/:imageId", h.DeleteImage())
}