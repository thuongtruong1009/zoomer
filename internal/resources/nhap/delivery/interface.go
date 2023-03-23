package delivery

import (
	"github.com/labstack/echo/v4"
)

type ResourceHandler interface {
	UploadImage() echo.HandlerFunc
	DownloadImage() echo.HandlerFunc
	DeleteImage() echo.HandlerFunc
}
