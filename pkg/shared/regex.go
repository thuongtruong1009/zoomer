package shared

import "regexp"

const (
	EmptyRegex = `^$`
	SpaceRegex = `\s+`
	WordNumRegex = `^[a-zA-Z0-9]+$`
)

const (
	UsernameLenRegex = `^.{4,20}$`
	PasswordLenRegex = `^.{8,32}$`
	EmailLenRegex    = `^.{8,32}$`
)

func MatchRegex(regex, chain string) bool {
	return regexp.MustCompile(regex).MatchString(chain)
}
