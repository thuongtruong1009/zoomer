package presenter

import "time"

type SetCookie struct {
	Name string
	Value string
	Expires time.Duration
}

type ForgotPassword struct {
	Email string `json:"email"`
}

type ResetPassword struct {
	Password string `json:"password"`
	NewPassword string `json:"new_password"`
}

func (input *ResetPassword) IsPasswordValid() error {
	return validatePassword(input.NewPassword)
}
