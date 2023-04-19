package middlewares

import (
	"zoomer/internal/auth/usecase"
	"zoomer/pkg/interceptor"
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
