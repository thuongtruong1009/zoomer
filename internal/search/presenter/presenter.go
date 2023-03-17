package presenter

import "zoomer/internal/models"

type RoomSearch struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

func (r *RoomSearch) ParseToModel() *models.RoomSearch {
	category := models.Category(r.Category)
	return &models.RoomSearch{
		Name:        r.Name,
		Description: r.Description,
		Category:    category,
	}
}
