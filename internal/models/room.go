package models

import "time"

type Room struct {
	Id          string `gorm:"primary_key"`
	Name        string
	Description string
	Category    Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   string
	User        User `gorm:"foreignKey:CreatedBy"`
}
