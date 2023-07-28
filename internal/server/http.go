package server

import (
	"github.com/labstack/echo/v4"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/middlewares"

	authRepository "github.com/thuongtruong1009/zoomer/internal/auth/repository"
	resourceRepository "github.com/thuongtruong1009/zoomer/internal/resources/minio/repository"
	roomRepository "github.com/thuongtruong1009/zoomer/internal/rooms/repository"
	searchRepository "github.com/thuongtruong1009/zoomer/internal/search/repository"

	authUsecase "github.com/thuongtruong1009/zoomer/internal/auth/usecase"
	localResourceUsecase "github.com/thuongtruong1009/zoomer/internal/resources/local/usecase"
	resourceUsecase "github.com/thuongtruong1009/zoomer/internal/resources/minio/usecase"
	roomUsecase "github.com/thuongtruong1009/zoomer/internal/rooms/usecase"
	searchUsecase "github.com/thuongtruong1009/zoomer/internal/search/usecase"

	authHttp "github.com/thuongtruong1009/zoomer/internal/auth/delivery"
	localResourceHttp "github.com/thuongtruong1009/zoomer/internal/resources/local/delivery"
	minioResourceHttp "github.com/thuongtruong1009/zoomer/internal/resources/minio/delivery"
	roomHttp "github.com/thuongtruong1009/zoomer/internal/rooms/delivery"
	searchHttp "github.com/thuongtruong1009/zoomer/internal/search/delivery"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/thuongtruong1009/zoomer/docs"
)

func (s *Server) HttpMapServer(e *echo.Echo) error {
	authRepo := authRepository.NewAuthRepository(s.pgDB, s.redisDB)
	roomRepo := roomRepository.NewRoomRepository(s.pgDB, s.redisDB)
	roomRepo.CreateFetchChatBetweenIndex()
	searchRepo := searchRepository.NewSearchRepository(s.pgDB)
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

	e.Static(constants.StaticGroupPath, constants.StaticGroupName)
	// e.Use(middleware.Static("/public"))
	e.GET(constants.DocGroup, echoSwagger.WrapHandler)

	middlewares.HttpMiddleware(e, s.inter)
	mw := middlewares.BaseMiddlewareManager(authUC, s.inter)

	httpGr := e.Group(constants.ApiGroup)

	authGroup := httpGr.Group(constants.AuthGroupEndPoint)
	authHttp.MapAuthRoutes(authGroup, authHandler)

	roomGroup := httpGr.Group(constants.RoomGroupEndPoint)
	roomHttp.MapRoomRoutes(roomGroup, roomHandler, mw)

	searchGroup := httpGr.Group(constants.SearchGroupEndPoint)
	searchHttp.MapSearchRoutes(searchGroup, searchHandler)

	resourceGroup := httpGr.Group(constants.ResourceGroupEndPoint)
	minioResourceHttp.MapResourceRoutes(resourceGroup, minioResourceHandler)
	localResourceHttp.MapLocalResourceRoutes(resourceGroup, localResourceHandler)

	return nil
}
