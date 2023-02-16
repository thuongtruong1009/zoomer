package presenter

type SignUpResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Limit    int    `json:"limit"`
}

type SignUpInput struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Limit    int    `json:"limit"`
}
