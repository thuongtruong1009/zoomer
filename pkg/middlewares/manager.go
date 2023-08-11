package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
)

type MiddlewareManager struct {
	e        *echo.Echo
	cfg      *configs.Configuration
	paramCfg *parameter.ParameterConfig
	inter    interceptor.IInterceptor
}

func RegisterMiddleware(e *echo.Echo, cfg *configs.Configuration, paramCfg *parameter.ParameterConfig, inter interceptor.IInterceptor) *MiddlewareManager {
	return &MiddlewareManager{
		e:        e,
		cfg:      cfg,
		paramCfg: paramCfg,
		inter:    inter,
	}
}
