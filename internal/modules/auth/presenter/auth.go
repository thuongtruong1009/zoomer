package presenter

import "time"

type MailRequest struct {
	Email string `json:"email"`
}

type VerifyOtp struct {
	Code string `json:"code"`
}

type SetCookie struct {
	Name    string
	Value   string
	Expires time.Duration
}

type UpdatePassword struct {
	Email 	 string `json:"email"`
	NewPassword string `json:"new_password"`
}

func (input *UpdatePassword) IsPasswordValid() error {
	return validatePassword(input.NewPassword)
}
