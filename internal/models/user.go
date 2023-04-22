package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       string `gorm:"primary_key"`
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Limit    int   `gorm:"not null" json:"limit"`
	CreateAt time.Time
	UpdateAt time.Time
}
//https://github.dev/dilaragorum/online-ticket-project-go/tree/master/internal/notification
const (
	MinPasswordLength int = 8
	MaxPasswordLength int = 20
)

func (u *User) IsNameEmpty() bool {
	return u.Username == ""
}

func (u *User) IsAuthTypeInvalid() bool {
	return !u.IsUserTypeValid()
}

func (u *User) IsUserTypeValid() bool {
	switch u.UserType {

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
