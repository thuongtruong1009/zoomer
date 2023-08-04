package server

import (
	"context"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/thuongtruong1009/zoomer/docs"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/middlewares"

	authRepository "github.com/thuongtruong1009/zoomer/internal/auth/repository"
	minioResourceRepository "github.com/thuongtruong1009/zoomer/internal/resources/minio/repository"
	roomRepository "github.com/thuongtruong1009/zoomer/internal/rooms/repository"
	searchRepository "github.com/thuongtruong1009/zoomer/internal/search/repository"

	authUsecase "github.com/thuongtruong1009/zoomer/internal/auth/usecase"
	localResourceUsecase "github.com/thuongtruong1009/zoomer/internal/resources/local/usecase"
	minioResourceUsecase "github.com/thuongtruong1009/zoomer/internal/resources/minio/usecase"
	roomUsecase "github.com/thuongtruong1009/zoomer/internal/rooms/usecase"
	searchUsecase "github.com/thuongtruong1009/zoomer/internal/search/usecase"

	authHttp "github.com/thuongtruong1009/zoomer/internal/auth/delivery"
	localResourceHttp "github.com/thuongtruong1009/zoomer/internal/resources/local/delivery"
	minioResourceHttp "github.com/thuongtruong1009/zoomer/internal/resources/minio/delivery"
	roomHttp "github.com/thuongtruong1009/zoomer/internal/rooms/delivery"
	searchHttp "github.com/thuongtruong1009/zoomer/internal/search/delivery"
)

type IServer interface {
	HttpMapServer() error
	WsMapServer() error
	start()
	Notify() <-chan error
	Shutdown(ctx context.Context) error
}

func (s *Server) HttpMapServer() error {
	authRepo := authRepository.NewAuthRepository(s.pgDB, s.redisDB)
	roomRepo := roomRepository.NewRoomRepository(s.pgDB, s.redisDB)

	roomRepo.CreateFetchChatBetweenIndex()

	searchRepo := searchRepository.NewSearchRepository(s.pgDB)
	minioRepository := minioResourceRepository.NewResourceRepository()

	authUC := authUsecase.NewAuthUseCase(authRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	roomUC := roomUsecase.NewRoomUseCase(roomRepo, authRepo)
	searchUC := searchUsecase.NewSearchUseCase(searchRepo, roomRepo)
	minioResourceUC := minioResourceUsecase.NewMinioResourceUseCase(s.minioClient, minioRepository)
	localResourceUC := localResourceUsecase.NewLocalResourceUseCase()

	authHandler := authHttp.NewAuthHandler(authUC, s.inter)
	roomHandler := roomHttp.NewRoomHandler(roomUC, s.inter)
	searchHandler := searchHttp.NewSearchHandler(searchUC)
	minioResourceHandler := minioResourceHttp.NewResourceHandler(minioResourceUC)
	localResourceHandler := localResourceHttp.NewLocalResourceHandler(s.inter, localResourceUC)

	s.echo.Static(constants.StaticGroupPath, constants.StaticGroupName)
	// e.Use(middleware.Static(constants.StaticGroupPath))
	s.echo.GET(constants.DocGroup, echoSwagger.WrapHandler)

	mw := middlewares.RegisterMiddleware(s.echo, s.cfg, s.parameterCfg, s.inter)
	mw.HttpMiddleware()
	authMw := middlewares.NewAuthMiddleware(authUC, s.inter)

	httpGr := s.echo.Group(constants.ApiGroup)

	authGroup := httpGr.Group(constants.AuthGroupEndPoint)
	authHttp.MapAuthRoutes(authGroup, authHandler, authMw)

	roomGroup := httpGr.Group(constants.RoomGroupEndPoint)
	roomHttp.MapRoomRoutes(roomGroup, roomHandler, authMw)

	searchGroup := httpGr.Group(constants.SearchGroupEndPoint)
	searchHttp.MapSearchRoutes(searchGroup, searchHandler)

	resourceGroup := httpGr.Group(constants.ResourceGroupEndPoint)

	minioResourceHttp.MapResourceRoutes(resourceGroup, minioResourceHandler)
	localResourceHttp.MapLocalResourceRoutes(resourceGroup, localResourceHandler)

	return nil
}
