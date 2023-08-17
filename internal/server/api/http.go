package api

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/thuongtruong1009/zoomer/docs"
	"github.com/thuongtruong1009/zoomer/internal/middlewares"
	"github.com/thuongtruong1009/zoomer/pkg/constants"

	"github.com/thuongtruong1009/zoomer/infrastructure/mail"

	authRepository "github.com/thuongtruong1009/zoomer/internal/modules/auth/repository"
	minioResourceRepository "github.com/thuongtruong1009/zoomer/internal/modules/resources/minio/repository"
	roomRepository "github.com/thuongtruong1009/zoomer/internal/modules/rooms/repository"
	searchRepository "github.com/thuongtruong1009/zoomer/internal/modules/search/repository"
	userRepository "github.com/thuongtruong1009/zoomer/internal/modules/users/repository"

	authUsecase "github.com/thuongtruong1009/zoomer/internal/modules/auth/usecase"
	localResourceUsecase "github.com/thuongtruong1009/zoomer/internal/modules/resources/local/usecase"
	minioResourceUsecase "github.com/thuongtruong1009/zoomer/internal/modules/resources/minio/usecase"
	roomUsecase "github.com/thuongtruong1009/zoomer/internal/modules/rooms/usecase"
	searchUsecase "github.com/thuongtruong1009/zoomer/internal/modules/search/usecase"
	userUsecase "github.com/thuongtruong1009/zoomer/internal/modules/users/usecase"

	authHttp "github.com/thuongtruong1009/zoomer/internal/modules/auth/delivery"
	localResourceHttp "github.com/thuongtruong1009/zoomer/internal/modules/resources/local/delivery"
	minioResourceHttp "github.com/thuongtruong1009/zoomer/internal/modules/resources/minio/delivery"
	roomHttp "github.com/thuongtruong1009/zoomer/internal/modules/rooms/delivery"
	searchHttp "github.com/thuongtruong1009/zoomer/internal/modules/search/delivery"
	userHttp "github.com/thuongtruong1009/zoomer/internal/modules/users/delivery"
)

func (s *Api) HttpApi() error {
	mailUC := mail.NewMail(s.Cfg)

	authRepo := authRepository.NewAuthRepository(s.PgDB, s.RedisDB, s.ParameterCfg)
	roomRepo := roomRepository.NewRoomRepository(s.PgDB, s.RedisDB)
	userRepo := userRepository.NewUserRepository(s.PgDB, s.RedisDB, s.ParameterCfg)
	roomRepo.CreateFetchChatBetweenIndex()
	searchRepo := searchRepository.NewSearchRepository(s.PgDB)
	minioRepository := minioResourceRepository.NewResourceRepository()

	authUC := authUsecase.NewAuthUseCase(authRepo, userRepo, s.Cfg, s.ParameterCfg, mailUC)
	roomUC := roomUsecase.NewRoomUseCase(roomRepo, userRepo)
	searchUC := searchUsecase.NewSearchUseCase(searchRepo, roomRepo)
	minioResourceUC := minioResourceUsecase.NewMinioResourceUseCase(s.MinioClient, minioRepository)
	localResourceUC := localResourceUsecase.NewLocalResourceUseCase()
	userUC := userUsecase.NewUserUseCase(userRepo)

	authHandler := authHttp.NewAuthHandler(authUC, s.Inter, s.ParameterCfg)
	roomHandler := roomHttp.NewRoomHandler(roomUC, s.Inter)
	searchHandler := searchHttp.NewSearchHandler(searchUC)
	minioResourceHandler := minioResourceHttp.NewResourceHandler(minioResourceUC)
	localResourceHandler := localResourceHttp.NewLocalResourceHandler(s.Inter, localResourceUC)
	userHandler := userHttp.NewUserHandler(userUC, s.Inter)

	s.Echo.Static(constants.StaticGroupPath, constants.StaticGroupName)
	// e.Use(middleware.Static(constants.StaticGroupPath))
	s.Echo.GET(constants.DocSpecial, echoSwagger.WrapHandler)
	// s.Echo.GET(constants.DocCommon, echoSwagger.WrapHandler)

	mw := middlewares.RegisterMiddleware(s.Echo, s.Cfg, s.ParameterCfg, s.Inter)
	mw.HttpMiddleware()
	authMw := middlewares.NewAuthMiddleware(authUC, s.Inter)

	httpGr := s.Echo.Group(constants.ApiGroup)

	authGroup := httpGr.Group(constants.AuthGroupEndPoint)
	authHttp.MapAuthRoutes(authGroup, authHandler, authMw)

	roomGroup := httpGr.Group(constants.RoomGroupEndPoint)
	roomHttp.MapRoomRoutes(roomGroup, roomHandler, authMw)

	searchGroup := httpGr.Group(constants.SearchGroupEndPoint)
	searchHttp.MapSearchRoutes(searchGroup, searchHandler)

	resourceGroup := httpGr.Group(constants.ResourceGroupEndPoint)

	minioResourceHttp.MapResourceRoutes(resourceGroup, minioResourceHandler)
	localResourceHttp.MapLocalResourceRoutes(resourceGroup, localResourceHandler)

	userGroup := httpGr.Group(constants.UserGroupEndPoint)
	userHttp.MapUserRoutes(userGroup, userHandler)

	return nil
}
