package middlewares

import (
	"zoomer/internal/auth"
)

type MiddlewareManager struct {
	authUC auth.UseCase
	// cfg    *config.Configuration
	// logger *logrus.Logger
	// origins []string
}

func NewMiddlewareManager(authUC auth.UseCase) *MiddlewareManager {
	return &MiddlewareManager{authUC}
}
