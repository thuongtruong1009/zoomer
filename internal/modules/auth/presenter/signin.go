package presenter

type SignInInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (input *SignInInput) IsUsernameValid() error {
	return validateUsername(input.Username)
}

func (input *SignInInput) IsPasswordValid() error {
	return validatePassword(input.Password)
}

type SignInResponse struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
