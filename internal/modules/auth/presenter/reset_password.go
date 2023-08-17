package presenter

type ResetPassword struct {
	Password string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" validate:"eqfield=Password"`
}
