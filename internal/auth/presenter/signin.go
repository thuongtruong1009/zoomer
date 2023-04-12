package presenter

type LoginInput struct {
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,password"`
}

type LogInResponse struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
