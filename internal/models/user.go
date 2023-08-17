package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       string `gorm:"primary_key"`
	Username string `gorm:"not null;unique;type:varchar(100)" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Limit    int    `gorm:"not null" json:"limit"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}
