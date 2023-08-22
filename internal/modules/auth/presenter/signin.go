package presenter

import "strings"

type SignInRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password string `json:"password"`
}

func (input *SignInRequest) IsRequestValid() error {
	iAt := strings.IndexByte(input.UsernameOrEmail, '@')
    if iAt <=2 || iAt > -2 {
        return validateEmail(input.UsernameOrEmail)
    }

	err := validateUsername(input.UsernameOrEmail)
	if err != nil {
		return err
	}

	err = validatePassword(input.Password)
	if err != nil {
		return err
	}

	return nil
}

type SignInResponse struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Email string `json:"email"`
	Token    string `json:"token"`
}
