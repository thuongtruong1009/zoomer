package presenter

import "time"

type ForgotPassword struct {
	Email string `json:"email"`
}

type VerifyResetPasswordOtp struct {
	Code string `json:"code"`
}

type SetCookie struct {
	Name    string
	Value   string
	Expires time.Duration
}

type ResetPassword struct {
	Email 	 string `json:"email"`
	NewPassword string `json:"new_password"`
}


func (input *ResetPassword) IsPasswordValid() error {
	return validatePassword(input.NewPassword)
}
