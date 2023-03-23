package presenter

type ResourceResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type ResourceRequest struct {
	Name string `json:"name"`
}
