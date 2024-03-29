package hub

import (
	"context"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/internal/models"
	"github.com/thuongtruong1009/zoomer/pkg/helpers"
	"log"
	"sync"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
)

type RoomMap struct {
	Map map[string][]*models.Participant
	mux sync.RWMutex
}

var (
	Mapper     RoomMap
	Broadcast  chan *models.BroadcastMessage
	Disconnect chan *models.DisconnectMessage
)

type hub struct{}

func NewStreamHub(streamCfg *parameter.ServerConf) IHub {
	Mapper = RoomMap{Map: make(map[string][]*models.Participant)}
	Broadcast = make(chan *models.BroadcastMessage, streamCfg.StreamMaxConnection)
	Disconnect = make(chan *models.DisconnectMessage, streamCfg.StreamMaxConnection)
	return &hub{}
}

func (h *hub) CreateStream(ctx context.Context) string {
	roomID := helpers.RandomChain(constants.RandomTypeString, 8)
	Mapper.mux.Lock()
	Mapper.Map[roomID] = []*models.Participant{}
	Mapper.mux.Unlock()

	return roomID
}

func (h *hub) GetParticipants(ctx context.Context, roomID string) []*models.Participant {
	return Mapper.Map[roomID]
}

func (h *hub) InsertIntoStream(ctx context.Context, roomID string, client *models.Participant) {
	Mapper.mux.Lock()
	defer Mapper.mux.Unlock()
	Mapper.Map[roomID] = append(Mapper.Map[roomID], client)
	log.Println("inserted into room", roomID, "client", client)
}

func (h *hub) DeleteStream(ctx context.Context, roomID string) {
	Mapper.mux.Lock()
	defer Mapper.mux.Unlock()
	delete(Mapper.Map, roomID)
}

func (h *hub) Receiver(ctx context.Context, roomId string, client *models.Participant) {
	for {
		var msg models.BroadcastMessage
		err := client.Conn.ReadJSON(&msg.Message)
		if err != nil {
			log.Println("Can read messaged: ", err)
			// Disconnect <- &models.DisconnectMessage{RoomID: roomId, Client: client}
			return
		}

		msg.Client = client.Conn
		msg.RoomID = roomId

		Broadcast <- &msg
	}
}

func (h *hub) Broadcaster() {
	for {
		select {
		case msg := <-Broadcast:
			Mapper.mux.RLock()
			participants, ok := Mapper.Map[msg.RoomID]
			Mapper.mux.RUnlock()
			if !ok {
				log.Println("room", msg.RoomID, "not found")
				continue
			}

			for _, client := range participants {
				if client.Conn != msg.Client {
					err := client.Conn.WriteJSON(msg.Message)
					if err != nil {
						log.Println("error while broadcasting message to client", client.Conn, err)
						// client.Conn.Close()
						// Mapper.Map = nil
						// log.Println("Room remain: ", Mapper.Map)
						// Disconnect <- &models.DisconnectMessage{RoomID: msg.RoomID, Client: client}
					}
				}
			}
		case msg := <-Disconnect:
			Mapper.mux.Lock()
			defer Mapper.mux.Unlock()
			roomID := msg.RoomID
			client := msg.Client
			clients := Mapper.Map[roomID]
			for i, c := range clients {
				if c.Conn == client.Conn {
					clients = append(clients[:i], clients[i+1:]...)
					Mapper.Map[roomID] = clients
					log.Println("client disconnected from room", roomID)
					// break
				}
			}
		}
	}
}
