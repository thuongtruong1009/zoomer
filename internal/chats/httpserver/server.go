package httpserver

import (
	"net/http"

	"zoomer/db"
	"zoomer/configs"
	"zoomer/internal/chats/repository"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartHTTPServer() {

	redisClient := db.GetRedisInstance(configs.NewConfig())
	defer redisClient.Close()

	// create indexes
	repository.CreateFetchChatBetweenIndex()

	r := mux.NewRouter()

	r.HandleFunc("/chat-history", chatHistoryHandler).Methods(http.MethodGet)
	r.HandleFunc("/contact-list", contactListHandler).Methods(http.MethodGet)

	// Use default options
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":8080", handler)
}
