package presenter

type LogInResponse struct {
	UserId string `json:"userId"`
	Username string `json:"username"`
	Token string `json:"token"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
