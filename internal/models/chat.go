package models

type Chat struct {
	ID        string    `gorm:"primary_key" json:"id"`
	From string `json:"from"`
	To string `json:"to"`
	Message string `json:"message"`
	Image string `json:"image" gorm:"default:null" required:"false"`
	Timestamp int64  `json:"timestamp"`
}

type ContactList struct {
	Username string `json:"username"`
	LastActivity int64 `json:"last_activity"`
}