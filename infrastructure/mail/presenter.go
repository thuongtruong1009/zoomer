package mail

type Mail struct {
	To string `json:"to"`
	Subject string `json:"subject"`
	Body string `json:"body"`
}
