package server

import (
	chatWs "zoomer/internal/chats"

	"github.com/labstack/echo/v4"
)

func WsMapHandlers() error {
	e := echo.New()

	wsUC := chatWs.NewHub()

	wsHandler := chatWs.NewChatHandler(wsUC)

	chatWs.MapChatRoutes(e, wsHandler, "/api/chats")

	e.Logger.Fatal(e.Start(":8081"))

	defer e.Close()
	go wsUC.Run()

	return nil
}
