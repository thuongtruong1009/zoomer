package models

import "time"

type Todo struct {
	Id        string    `gorm:"primary_key" json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	User      User      `gorm:"foreignKey:CreatedBy" json:"user"`
}
