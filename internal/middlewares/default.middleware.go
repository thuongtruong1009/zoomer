package middlewares

import (
	"zoomer/internal/auth/usecase"
)

type MiddlewareManager struct {
	authUC usecase.UseCase
	// cfg    *config.Configuration
	// logger *logrus.Logger
	// origins []string
}

func NewMiddlewareManager(authUC usecase.UseCase) *MiddlewareManager {
	return &MiddlewareManager{authUC}
}
