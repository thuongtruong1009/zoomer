package presenter

type SignUpRequest struct {
	Username string `json:"username" validate:"required,min=4,max=20"`
	Email string `json:"email" validate:"required,min=12,max=32"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Limit    int    `json:"limit" validate:"required,max=10"`
}

func (input *SignUpRequest) IsRequestValid() error {
	err := validateUsername(input.Username)
	if err != nil {
		return err
	}

	err = validateEmail(input.Email)
	if err != nil {
		return err
	}

	err = validatePassword(input.Password)
	if err != nil {
		return err
	}

	err = validateLimit(input.Limit)
	if err != nil {
		return err
	}

	return nil
}

type SignUpResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Limit    int    `json:"limit"`
}
