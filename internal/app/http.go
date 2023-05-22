package app

import (
	"github.com/labstack/echo/v4"
	"zoomer/pkg/middlewares"

	authRepository "zoomer/internal/auth/repository"
	roomRepository "zoomer/internal/rooms/repository"
	searchRepository "zoomer/internal/search/repository"
	resourceRepository "zoomer/internal/resources/minio/repository"

	authUsecase "zoomer/internal/auth/usecase"
	roomUsecase "zoomer/internal/rooms/usecase"
	searchUsecase "zoomer/internal/search/usecase"
	resourceUsecase "zoomer/internal/resources/minio/usecase"
	localResourceUsecase "zoomer/internal/resources/local/usecase"

	authHttp "zoomer/internal/auth/delivery"
	roomHttp "zoomer/internal/rooms/delivery"
	searchHttp "zoomer/internal/search/delivery"
	minioResourceHttp "zoomer/internal/resources/minio/delivery"
	localResourceHttp "zoomer/internal/resources/local/delivery"

	echoSwagger "github.com/swaggo/echo-swagger"
	"zoomer/pkg/interceptor"
)

func (s *Server) HttpMapServer(e *echo.Echo) error {
	inter := interceptor.NewInterceptor()

	userRepo := authRepository.NewUserRepository(s.db)
	roomRepo := roomRepository.NewRoomRepository(s.db)
	searchRepo := searchRepository.NewSearchRepository(s.db)
	resourceRepository := resourceRepository.NewResourceRepository()

	authUC := authUsecase.NewAuthUseCase(userRepo, s.cfg.HashSalt, []byte(s.cfg.SigningKey), s.cfg.TokenTTL)
	roomUC := roomUsecase.NewRoomUseCase(roomRepo, userRepo)
	searchUC := searchUsecase.NewSearchUseCase(searchRepo, roomRepo)
	resourceUC := resourceUsecase.NewResourceUseCase(resourceRepository)
	localResourceUC := localResourceUsecase.NewLocalResourceUseCase()

	authHandler := authHttp.NewAuthHandler(authUC, inter)
	roomHandler := roomHttp.NewRoomHandler(roomUC, inter)
	searchHandler := searchHttp.NewSearchHandler(searchUC)
	minioResourceHandler := minioResourceHttp.NewResourceHandler(resourceUC)
	localResourceHandler := localResourceHttp.NewLocalResourceHandler(inter, localResourceUC)

	middlewares.HttpMiddleware(e, inter)

	mw := middlewares.BaseMiddlewareManager(authUC, inter)

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
