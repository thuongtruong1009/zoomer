package presenter

import (
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"unicode"
	// "strings"
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

func validateEmail(email string) error {
    if email == "" {
        return constants.ErrRequiredEmail
    }

	for _, char := range email {
        if unicode.IsSpace(char) {
            return constants.ErrSpaceEmail
        }
    }

    // iAt := strings.IndexByte(email, '@')
    // if iAt <=2 || iAt > -2 {
    //     return constants.ErrInvalidEmail
    // }

    // if email[:iAt] == "" {
    //     return constants.ErrInvalidEmail
    // }

    // domain := email[iAt+1:]
    // if domain == "" {
    //     return constants.ErrInvalidEmail
    // }

    // iDot := strings.IndexByte(domain, '.')
    // if iDot == -1 || iDot == 0 {
    //     return constants.ErrInvalidEmail
    // }

    // if strings.Index(domain, "..") != -1 {
    //     return constants.ErrInvalidEmail
    // }

    // iTLD := strings.LastIndexByte(domain, '.')
    // if len([]rune(domain[iTLD+1:])) >= 2 {
	// 	return constants.ErrInvalidEmail
	// }
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
