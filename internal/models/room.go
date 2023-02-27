package models

import "time"

type Room struct {
	Id        string    `gorm:"primary_key"`
	Name      string
	CreatedAt time.Time
	CreatedBy string
	User      User      `gorm:"foreignKey:CreatedBy"`
}
