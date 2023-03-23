package models

type ResourceList struct {
	ResourceList []Resource `json:"resource"`
}
type Resource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}