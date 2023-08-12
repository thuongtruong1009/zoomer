package presenter

type SearchResquest struct {
	Username string `json:"search"`
}

type SearchResponse struct {
	Match []SignInResponse
}
