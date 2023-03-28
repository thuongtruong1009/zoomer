package httpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"zoomer/internal/chats/repository"
)

type userReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Client   string `json:"client"`
}

type response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   int         `json:"total,omitempty"`
}