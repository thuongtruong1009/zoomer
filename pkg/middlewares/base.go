package middlewares

import (
	"github.com/thuongtruong1009/zoomer/internal/auth/usecase"
	"github.com/thuongtruong1009/zoomer/pkg/interceptor"
)

type MiddlewareManager struct {
	authUC usecase.UseCase
	inter  interceptor.IInterceptor
	// cfg    *config.Configuration
	// logger *logrus.Logger
	// origins []string
}

func BaseMiddlewareManager(authUC usecase.UseCase, inter interceptor.IInterceptor) *MiddlewareManager {
	return &MiddlewareManager{authUC, inter}
}
