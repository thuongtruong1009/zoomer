package models

import (
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"golang.org/x/crypto/bcrypt"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
)

type User struct {
	Id       string `gorm:"primary_key"`
	Username string `gorm:"not null;unique;type:varchar(20)" json:"username"`
	Email string `gorm:"not null;unique;type:varchar(32)" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Limit    int    `gorm:"not null" json:"limit"`
}

type UsersList struct {
	TotalCount int64    `json:"total_count"`
	TotalPages int64    `json:"total_pages"`
	Page       int64    `json:"page"`
	Size       int64    `json:"size"`
	HasMore    bool     `json:"has_more"`
	Users     []*User `json:"users"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		exceptions.Log(constants.ErrHashPassword, err)
		return constants.ErrHashPassword
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		exceptions.Log(constants.ErrComparePassword, err)
		return constants.ErrComparePassword
	}

	return nil
}
