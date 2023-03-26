package server

import (
	"github.com/labstack/echo/v4"
	chatWs "zoomer/internal/chats"
	middlewares "zoomer/internal/middlewares"
)

func WsMapHandlers(port string) {
	e := echo.New()
	defer e.Close()

	e.Use(middlewares.WsCORS)
	wsUC := chatWs.NewHub()
	go wsUC.Run()

	wsHandler := chatWs.NewChatHandler(wsUC)

	chatWs.MapChatRoutes(e, wsHandler, "/api/chats")

	e.Logger.Fatal(e.Start(port))
}
