package views

import (
	"zoomer/internal/models"
)

type RoomFind struct {
	Name       string    `json:"name"`
	Description string    `json:"desc"`
	Category    string    `json:"category"`
}

func NewRoomsFind(roomsModel *[]models.Room) *[]RoomFind {
	var rooms []RoomFind

	for _, room := range *roomsModel {
		var tempRoom = RoomFind{
			Name:        room.Name,
			Description: room.Description,
			Category:    room.Category.String(),
		}
		rooms = append(rooms, tempRoom)
	}
	return &rooms
}
