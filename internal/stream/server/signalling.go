package stream

import (
  "net/http"
  "log"
  "github.com/gorilla/websocket"
  "encoding/json"
)

var AllRooms RoomMap

func CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  roomID := AllRooms.CreateRoom()

  type resp struct {
    RoomID string `json:"room_id"`
  }

  log.Println(AllRooms.Map)
  json.NewEncoder(w).Encode(resp{RoomID: roomID})
}

var upgrader = websocket.Upgrader {
  ReadBufferSize: 1024,
  WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool {
    return true
  },
}

type broadcastMsg struct {
  Message map[string]interface{} `json:"message"`
  RoomID string `json:"room_id"`
  Client *websocket.Conn `json:"conn"`
}

var broadcast = make(chan broadcastMsg)

func broadcaster(){
  for {
    msg := <- broadcast

    for _, client := range AllRooms.Map[msg.RoomID] {
      if (client.Conn != msg.Client) {
        err := client.Conn.WriteJSON(msg.Message)

        if err != nil {
          log.Fatal(err)
          client.Conn.Close()
        }
      }
    }
  }
}

func JoinRoomRequestHandler(w http.ResponseWriter, r *http.Request) {
  roomID, ok := r.URL.Query()["roomID"]

  if !ok || len(roomID[0]) < 1 {
    log.Println("Url Param 'roomID' is missing")
    return
  }

  // ws, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
  ws, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Fatal("Web socket upgrade failed: ", err)
  }

  AllRooms.InsertIntoRoom(roomID[0], false, ws)

  go broadcaster()

  for {
    var msg broadcastMsg
    err := ws.ReadJSON(&msg.Message)
    if err != nil {
      log.Fatal("Read failed: ", err)
    }

    msg.Client = ws
    msg.RoomID = roomID[0]

    broadcast <- msg
  }
}
