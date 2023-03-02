package presenter

import "time"

type RoomResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
}

type RoomRequest struct {
	Name string `json:"name"`
}
