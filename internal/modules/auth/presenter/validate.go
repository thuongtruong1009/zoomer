package presenter

import (
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/shared"
	"unicode"
)

func validateUsername(username string) error {
	switch {
	case shared.MatchRegex(shared.EmptyRegex, username):
		return constants.ErrReqiredUsername
	case !shared.MatchRegex(shared.UsernameLenRegex, username):
		return constants.ErrLenUsername
	case shared.MatchRegex(shared.SpaceRegex, username):
		return constants.ErrSpaceUsername
	case !shared.MatchRegex(shared.WordNumRegex, username):
		return constants.ErrAlphaNumUsername
	}

	return nil
}

func validateEmail(email string) error {
	switch {
	case shared.MatchRegex(shared.EmptyRegex, email):
		return constants.ErrRequiredEmail
	case !shared.MatchRegex(shared.EmailLenRegex, email):
		return constants.ErrLenEmail
	case shared.MatchRegex(shared.SpaceRegex, email):
		return constants.ErrSpaceEmail

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
	switch {
	case shared.MatchRegex(shared.EmptyRegex, password):
		return constants.ErrRequiredPassword
	case !shared.MatchRegex(shared.PasswordLenRegex, password):
		return constants.ErrLenPassword
	case shared.MatchRegex(shared.SpaceRegex, password):
		return constants.ErrSpacePassword
	}

	var hasUpperCase, hasLowerCase, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
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
