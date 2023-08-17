package presenter

import "time"

type SetCookie struct {
	Name string
	Value string
	Expires time.Duration
}
