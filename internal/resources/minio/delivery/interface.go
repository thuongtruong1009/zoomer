package delivery

import "github.com/labstack/echo/v4"

type ResourceHandler interface {
	GetResource() echo.HandlerFunc

	CreateResource() echo.HandlerFunc

	UploadResource() echo.HandlerFunc

	DeleteResource() echo.HandlerFunc
}
