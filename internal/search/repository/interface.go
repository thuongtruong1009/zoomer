package repository

import "zoomer/internal/models"

type SearchRepository interface {
	FindRoomBySearch(search *models.RoomSearch) ([]*models.Room, error)
}
