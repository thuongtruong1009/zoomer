package presenter

type GetUserByIdOrNameResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Limit    int    `json:"limit"`
}
