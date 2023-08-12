package presenter

type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInResponse struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
