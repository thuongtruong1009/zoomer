package presenter

type GetUserByIdOrNameResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Limit    int    `json:"limit"`
}
