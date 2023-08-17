package presenter

import (
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"unicode"
)

func validateUsername(username string) error {
	if username == "" {
		return constants.ErrReqiredUsername
	}

	usernameLen := len(username)
	if usernameLen < constants.MinUsernameLen || usernameLen > constants.MaxUsernameLen {
		return constants.ErrLenUsername
	}

	for _, char := range username {
		if unicode.IsSpace(char){
			return constants.ErrSpaceUsername
		}
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return constants.ErrAlphaNumUsername
		}
	}

	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return constants.ErrReqiredPassword
	}

	passLen := len(password)
	if passLen < constants.MinPasswordLen || passLen > constants.MaxPasswordLen {
		return constants.ErrLenPassword
	}

	var hasUpperCase, hasLowerCase, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsSpace(char):
			return constants.ErrSpacePassword
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsUpper(char):
			hasUpperCase = true
		case unicode.IsLetter(char):
			hasLowerCase = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		default:
			hasUpperCase, hasLowerCase, hasNumber, hasSpecial = false, false, false, false
		}
	}

	if !hasUpperCase || !hasLowerCase || !hasNumber || !hasSpecial {
		return constants.ErrAlphaNumPassword
	}

	return nil
}

func validateLimit(limit int) error {
	if limit < 0 {
		return constants.ErrInvalidRoomLimit
	} else if limit > constants.MaxRoomLimit {
		return constants.ErrMaxRoomLimit
	}

	return nil
}
