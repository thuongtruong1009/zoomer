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

func chatHistoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// user1 user2
	u1 := r.URL.Query().Get("u1")
	u2 := r.URL.Query().Get("u2")

	// chat between timerange fromTS toTS
	// where TS is timestamp
	// 0 to positive infinity
	fromTS, toTS := "0", "+inf"

	if r.URL.Query().Get("from-ts") != "" && r.URL.Query().Get("to-ts") != "" {
		fromTS = r.URL.Query().Get("from-ts")
		toTS = r.URL.Query().Get("to-ts")
	}

	res := chatHistory(u1, u2, fromTS, toTS)
	json.NewEncoder(w).Encode(res)
}

func contactListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	u := r.URL.Query().Get("username")

	res := contactList(u)
	json.NewEncoder(w).Encode(res)
}

func chatHistory(username1, username2, fromTS, toTS string) *response {
	// if invalid usernames return error
	// if valid users fetch chats
	res := &response{}

	fmt.Println(username1, username2)
	// check if user exists
	if !repository.IsUserExist(username1) || !repository.IsUserExist(username2) {
		res.Message = "incorrect username"
		return res
	}

	chats, err := repository.FetchChatBetween(username1, username2, fromTS, toTS)
	if err != nil {
		log.Println("error in fetch chat between", err)
		res.Message = "unable to fetch chat history. please try again later."
		return res
	}

	res.Status = true
	res.Data = chats
	res.Total = len(chats)
	return res
}

func contactList(username string) *response {
	// if invalid username return error
	// if valid users fetch chats
	res := &response{}

	// check if user exists
	if !repository.IsUserExist(username) {
		res.Message = "incorrect username"
		return res
	}

	contactList, err := repository.FetchContactList(username)
	if err != nil {
		log.Println("error in fetch contact list of username: ", username, err)
		res.Message = "unable to fetch contact list. please try again later."
		return res
	}

	res.Status = true
	res.Data = contactList
	res.Total = len(contactList)
	return res
}
