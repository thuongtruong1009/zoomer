package views

import (
	"zoomer/internal/models"
	"github.com/google/uuid"
)

type RoomFind struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	ImgCover    string    `json:"img_cover"`
	Description string    `json:"desc"`
	Price       int       `json:"price"`
	Category    string    `json:"category"`
}

func NewRoomsFind(roomsModel *[]models.Room) *[]RoomFind {
	var rooms []RoomFind

	for _, room := range *roomsModel {
		var tempRoom = RoomFind{
			ID:          room.ID,
			Title:       room.Title,
			ImgCover:    room.ImgCover,
			Description: room.Description,
			Price:       room.Price,
			Category:    room.Category.String(),
		}
		rooms = append(rooms, tempRoom)
	}
	return &rooms
}
