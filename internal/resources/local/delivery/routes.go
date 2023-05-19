package delivery

import (
	"github.com/labstack/echo/v4"
)

func MapLocalResourceRoutes(localGroup *echo.Group, lh LocalHandler) {
	localGroup.POST("/local/single", lh.UploadSingleFile())
	localGroup.POST("/local/multiple", lh.UploadMultipleFile())

	localGroup.DELETE("/local/single/:fileName", lh.DeleteSingleFile())
	localGroup.DELETE("/local/multiple", lh.DeleteMultipleFile())
}
