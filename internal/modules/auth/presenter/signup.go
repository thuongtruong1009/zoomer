package presenter

type SignUpInput struct {
	Username string `json:"username" validate:"required,min=4,max=20"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Limit    int    `json:"limit" validate:"required,max=10"`
}

func (input *SignUpInput) IsUsernameValid() error {
	return validateUsername(input.Username)
}

func (input *SignUpInput) IsPasswordValid() error {
	return validatePassword(input.Password)
}

func (input *SignUpInput) IsLimitValid() error {
	return validateLimit(input.Limit)
}

type SignUpResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Limit    int    `json:"limit"`
}
