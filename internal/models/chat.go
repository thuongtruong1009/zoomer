package models

type Chat struct {
	ID string `json:"id"`
	From string `json:"from"`
	To string `json:"to"`
	Msg string `json:"msg"`
	MsgType string `json:"msg_type"`
	Timestamp int64 `json:"timestamp"`
}

type ContactList struct {
	Username string `json:"username"`
	LastActivity int64 `json:"last_activity"`
}