package models

import (
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `gorm:"primary_key"`
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Limit    int    `gorm:"not null" json:"limit"`
}

func (u *User) IsUsernameInvalid() bool {
	return u.Username != ""
}

func (u *User) IsPasswordInvalid() bool {
	passLength := len(u.Password)
	return passLength < constants.MinPasswordLen && passLength > constants.MaxPasswordLen
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}
