package presenter

type SignUpRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Limit    int    `json:"limit"`
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
