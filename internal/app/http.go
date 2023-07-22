package app

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/middlewares"

	authRepository "github.com/thuongtruong1009/zoomer/internal/auth/repository"
	roomRepository "github.com/thuongtruong1009/zoomer/internal/rooms/repository"
	searchRepository "github.com/thuongtruong1009/zoomer/internal/search/repository"
	resourceRepository "github.com/thuongtruong1009/zoomer/internal/resources/minio/repository"

	authUsecase "github.com/thuongtruong1009/zoomer/internal/auth/usecase"
	roomUsecase "github.com/thuongtruong1009/zoomer/internal/rooms/usecase"
	searchUsecase "github.com/thuongtruong1009/zoomer/internal/search/usecase"
	resourceUsecase "github.com/thuongtruong1009/zoomer/internal/resources/minio/usecase"
	localResourceUsecase "github.com/thuongtruong1009/zoomer/internal/resources/local/usecase"

	authHttp "github.com/thuongtruong1009/zoomer/internal/auth/delivery"
	roomHttp "github.com/thuongtruong1009/zoomer/internal/rooms/delivery"
	searchHttp "github.com/thuongtruong1009/zoomer/internal/search/delivery"
	minioResourceHttp "github.com/thuongtruong1009/zoomer/internal/resources/minio/delivery"
	localResourceHttp "github.com/thuongtruong1009/zoomer/internal/resources/local/delivery"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *Server) HttpMapServer(e *echo.Echo) error {
	pgInstance := s.pgDB.ConnectInstance(s.cfg)
	authRepo := authRepository.NewAuthRepository(pgInstance, s.redisDB)
	roomRepo := roomRepository.NewRoomRepository(pgInstance, s.redisDB)
	roomRepo.CreateFetchChatBetweenIndex()
	searchRepo := searchRepository.NewSearchRepository(pgInstance)
	resourceRepository := resourceRepository.NewResourceRepository()

	authUC := authUsecase.NewAuthUseCase(authRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	roomUC := roomUsecase.NewRoomUseCase(roomRepo, authRepo)
	searchUC := searchUsecase.NewSearchUseCase(searchRepo, roomRepo)
	resourceUC := resourceUsecase.NewResourceUseCase(resourceRepository)
	localResourceUC := localResourceUsecase.NewLocalResourceUseCase()

	authHandler := authHttp.NewAuthHandler(authUC, s.inter)
	roomHandler := roomHttp.NewRoomHandler(roomUC, s.inter)
	searchHandler := searchHttp.NewSearchHandler(searchUC)
	minioResourceHandler := minioResourceHttp.NewResourceHandler(resourceUC)
	localResourceHandler := localResourceHttp.NewLocalResourceHandler(s.inter, localResourceUC)

	middlewares.HttpMiddleware(e, s.inter)

	mw := middlewares.BaseMiddlewareManager(authUC, s.inter)

	e.Static("/", "public")

	httpGr := e.Group("/api")

	e.GET("/docs/*", echoSwagger.WrapHandler)
	/*
		url := echoSwagger.URL("http://localhost:1323/swagger/doc.json") //The url pointing to API definition
		e.GET("/docs/*", echoSwagger.EchoWrapHandler(url))
	*/

	authGroup := httpGr.Group("/auth")
	authHttp.MapAuthRoutes(authGroup, authHandler)

	roomGroup := httpGr.Group("/rooms")
	roomHttp.MapRoomRoutes(roomGroup, roomHandler, mw)

	searchGroup := httpGr.Group("/search")
	searchHttp.MapSearchRoutes(searchGroup, searchHandler)

	resourceGroup := httpGr.Group("/resource")
	minioResourceHttp.MapResourceRoutes(resourceGroup, minioResourceHandler)
	localResourceHttp.MapLocalResourceRoutes(resourceGroup, localResourceHandler)

	return nil
}
