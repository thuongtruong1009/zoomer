package models

import (
	"fmt"
)

type Category string

type RoomSearch struct {
	Name        string
	Category    Category
	Description string
}

const (
	Category_Unknown Category = "unknown"
	Category_Friend  Category = "friend"
	Category_Family  Category = "family"
	Category_Work    Category = "work"
	Category_School  Category = "school"
)

func (c Category) Validate() (err error) {
	switch c {
	case Category_Unknown, Category_Friend, Category_Family, Category_Work, Category_School:
		return
	default:
		err = fmt.Errorf("category: %v not found", c)
		return
	}
}

func (c Category) String() string {
	return string(c)
}
