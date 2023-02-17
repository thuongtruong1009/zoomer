package presenter

type MemberResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type MemberInput struct {
	Name string `json:"name"`
}
