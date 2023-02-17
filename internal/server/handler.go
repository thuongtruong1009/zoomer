package server

import (
	authRepository "zoomer/internal/auth/repository"
	authUseCase "zoomer/internal/auth/usecase"

	"zoomer/internal/middlewares"

	authHttp "zoomer/internal/auth/delivery"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) MapHandlers(e *echo.Echo) error [
	userRepo := authRepository.NewUserRepository(s.db)

	authUC := authUseCase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)

	authHandler := authHttp.NewAuthHandler(authUC)

	mw := middlewares.NewModdlewareManager(aurthUC)

	e.Use(middleware.BodyLimit("2M"))
	e.Use(mw.JWTValidation())
	e.Use(middleware.Logger())

	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")

	authHttp.MapAuthRoutes(authGroup, authHandler)

	return nil
]