package models

type Chat struct {
	ID string `json:"id"`
	From string `json:"from" validate:"required,from"`
	To string `json:"to" validate:"required,to"`
	Msg string `json:"msg" validate:"required,msg"`
	MsgType string `json:"msg_type" validate:"required,msg_type"`
	Timestamp int64 `json:"timestamp"`
}

type ContactList struct {
	Username string `json:"username"`
	LastActivity int64 `json:"last_activity"`
}