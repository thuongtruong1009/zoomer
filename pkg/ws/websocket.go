package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"zoomer/internal/models"
	"zoomer/pkg/redisrepo"
	"zoomer/db"
	"zoomer/configs"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Username string
}

type Message struct {
	Type string `json:"type"`
	User string `json:"user,omitempty"`
	Chat models.Chat `json:"chat,omitempty"`
}

var clients = make(map[*Client]bool)
var broadcast = make(chan *models.Chat)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true},
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host, r.URL.Query())
	ws, err := upgrader.Upgrade(w, r, nil) 
	if err != nil {
		log.Println(err)
	}
	

	client := &Client{Conn: ws}

	clients[client] = true
	fmt.Println("clients", len(clients), clients, ws.RemoteAddr())

	// listent new message
	receiver(client)

	fmt.Println("exiting", ws.RemoteAddr().String())
	delete(clients,client)
}

func receiver(client *Client) {
	for {
		_, p, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		m := &Message{}

		err = json.Unmarshal(p,m)
		if err != nil {
			log.Println("error while unmarshal chat", err)
			continue
		}

		fmt.Println("host", client.Conn.RemoteAddr())
		if m.Type == "bootup" {
			client.Username = m.User
			fmt.Println("clien successfully mapped", &client, client,client.Username)
		} else{
			fmt.Println("received message",m.Type, m.Chat)
			c := m.Chat
			c.Timestamp = time.Now().Unix()

			// save chat to redis
			id, err := redisrepo.CreateChat(&c)
			if err != nil {
				log.Println("error while saving chat", err)
				return
			}
			c.ID = id
			broadcast <- &c
		}
	}
}

func broadcaster() {
	for {
		message := <-broadcast
		fmt.Println("new message", message)

		for client := range clients {
			fmt.Println("username:",client.Username, "from:", message.From,"to:", message.To)

			if client.Username == message.From || client.Username == message.To{
				err:= client.Conn.WriteJSON(message)
				if err != nil {
					log.Printf("error: %v", err)
					client.Conn.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "chat server")
	})
	http.HandleFunc("/chat/ws", serveWs)
}

func StartWebsocketServer() {
	cfg := configs.NewConfig()	
	redisClient := db.InitialiseRedis(cfg)
	defer redisClient.Close()

	go broadcaster()
	setupRoutes()
	http.ListenAndServe(":8081", nil)
}
