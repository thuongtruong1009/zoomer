package models

import "time"

type ResourceList struct {
	ResourceList []Resource `json:"resource"`
}
type Resource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// video
type Video struct {
	Id        string
	Name      string
	Url       string
	CreatedAt time.Time
}
