package models

type User struct {
	Id       string `gorm:"primary_key"`
	Username string `gorm:"not null;unique;type:varchar(20)" json:"username"`
	Email    string `gorm:"not null;unique;type:varchar(32)" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Limit    int    `gorm:"not null" json:"limit"`
}

type UsersList struct {
	TotalCount int64   `json:"total_count"`
	TotalPages int64   `json:"total_pages"`
	Page       int64   `json:"page"`
	Size       int64   `json:"size"`
	HasMore    bool    `json:"has_more"`
	Users      []*User `json:"users"`
}
