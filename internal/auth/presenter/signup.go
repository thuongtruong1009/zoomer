package presenter

type SignUpInput struct {
	Username string `json:"username" validate:"required,username,min=3,max=20,unique"`
	Password string `json:"password" validate:"required,password,min=8,max=20"`
	Limit    int    `json:"limit"`
}

type SignUpResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Limit    int    `json:"limit"`
}
