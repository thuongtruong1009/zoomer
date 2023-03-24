package server

import (
	chatWs "zoomer/internal/chats"

	"github.com/labstack/echo/v4"
	middlewares "zoomer/internal/middlewares"
)

func WsMapHandlers(port string) error {
	e := echo.New()
	e.Use(middlewares.WsCORS)
	wsUC := chatWs.NewHub()

	wsHandler := chatWs.NewChatHandler(wsUC)

	chatWs.MapChatRoutes(e, wsHandler, "/api/chats")

	e.Logger.Fatal(e.Start(port))

	defer e.Close()
	go wsUC.Run()

	return nil
}
