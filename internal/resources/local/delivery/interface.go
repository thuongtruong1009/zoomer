package delivery

import (
	"github.com/labstack/echo/v4"
)

type LocalHandler interface {
	UploadSingleFile() echo.HandlerFunc
	UploadMultipleFile() echo.HandlerFunc
	DeleteSingleFile() echo.HandlerFunc
	DeleteMultipleFile() echo.HandlerFunc
}
