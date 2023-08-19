package presenter

type ResetPassword struct {
	Password string `json:"password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=20"`
}

func (input *ResetPassword) IsPasswordValid() error {
	return validatePassword(input.NewPassword)
}
