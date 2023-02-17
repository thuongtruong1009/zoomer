package models

import "time"

type Room struct {
	Id        string    `gorm:"primary_key" json:"id"`
	Name      string    `json:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	User      User      `gorm:"foreignKey:CreatedBy" json:"user"`
}
