package presenter

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogInResponse struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
