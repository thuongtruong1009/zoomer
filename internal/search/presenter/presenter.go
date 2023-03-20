package presenter

import "zoomer/internal/models"

type RoomSearchParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

func (r *RoomSearchParams) ParseToModel() *models.RoomSearch {
	category := models.Category(r.Category)
	return &models.RoomSearch{
		Name:        r.Name,
		Description: r.Description,
		Category:    category,
	}
}
