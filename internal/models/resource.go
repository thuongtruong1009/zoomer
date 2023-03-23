package models

type ResourceList struct {
	Resource []Resource `json:"resource"`
}
type Resource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}