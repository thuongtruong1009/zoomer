package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

func MapLocalResourceRoutes(localGroup *echo.Group, lh LocalHandler) {
	localGroup.POST(constants.UploadSingleLocalResourceEndPoint, lh.UploadSingleFile())
	localGroup.POST(constants.UploadMultipleLocalResourceEndPoint, lh.UploadMultipleFile())

	localGroup.DELETE(constants.DeleteSingleLocalResourceEndPoint, lh.DeleteSingleFile())
	localGroup.DELETE(constants.DeleteMultipleLocalResourceEndPoint, lh.DeleteMultipleFile())
}
