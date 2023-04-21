package hub

import (
	"context"
	// "encoding/json"
	"log"
	"math/rand"
	"time"
	"zoomer/internal/models"
)

type RoomMap struct {
	Map map[string][]*models.Participant
}

var (
	Mapper  RoomMap
	Broadcast = make(chan *models.BroadcastMessage, 100)
)

type hub struct{}

func NewStreamHub() IHub {
	return &hub{}
}

func (h *hub) CreateStream(ctx context.Context) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)
	Mapper.Map[roomID] = []*models.Participant{}

	return roomID
}

func (h *hub) GetParticipants(ctx context.Context, roomID string) []*models.Participant {
	return Mapper.Map[roomID]
}

func (h *hub) InsertIntoStream(ctx context.Context, roomID string, client *models.Participant) {
	Mapper.Map[roomID] = append(Mapper.Map[roomID], client)
	log.Println("inserted into room", roomID, "client", client)
}

func (h *hub) DeleteStream(ctx context.Context, roomID string) {
	delete(Mapper.Map, roomID)
}

func (h *hub) Receiver(ctx context.Context, roomId string, client *models.Participant) {
	for {
		// _, p, err := client.Conn.ReadMessage()
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }

		// m := &models.BroadcastMessage{}

		// err = json.Unmarshal(p, m)
		// if err != nil {
		// 	log.Println("error while unmarshaling stream broadcast", err)
		// 	continue
		// }

		// m.Client = client.Conn
		// m.RoomID = roomId

		// Broadcast <- m

		var msg models.BroadcastMessage
		err := client.Conn.ReadJSON(&msg.Message)
		if err != nil {
			log.Fatal("Read failed: ", err)
		}

		msg.Client = client.Conn
		msg.RoomID = roomId

		Broadcast <- &msg
	}
}

func (h *hub) Broadcaster() {
	for {
		msg := <-Broadcast

		participants, ok := Mapper.Map[msg.RoomID]
		if !ok {
			log.Println("room", msg.RoomID, "not found")
			continue
		}
		for _, client := range participants {
			if client.Conn != msg.Client {
				err := client.Conn.WriteJSON(msg.Message)
				if err != nil {
					log.Println("error while broadcasting message to client", client.Conn, err)
					client.Conn.Close()
				}
			}
		}
	}
}
