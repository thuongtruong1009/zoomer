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

//sync to redis
// type userReq struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Client   string `json:"client"`
// }

type ChatResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Total   int         `json:"total,omitempty"`
}
