package delivery

import "github.com/labstack/echo/v4"

func MapResourceRoutes(resourceGroup *echo.Group, rh ResourceHandler) {
	resourceGroup.GET("/image", rh.GetResource())
	resourceGroup.POST("/image/:uid/:id", rh.CreateResource())
	resourceGroup.PUT("/image/:uid/:id", rh.UploadResource())
	resourceGroup.DELETE("/image/:uid/:id", rh.DeleteResource())
	// resourceGroup.Logger.Fatal(router.Start(port))
}
