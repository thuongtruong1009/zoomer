package presenter

type RoomResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RoomInput struct {
	Name string `json:"name"`
}
