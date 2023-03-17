package repository

import (
	"errors"
	"gorm.io/gorm"
	"strings"
	"context"
	"zoomer/internal/models"
	"zoomer/internal/search/views"
)

type searchRepository struct {
	db *gorm.DB
}

func NewSearchRepository(db *gorm.DB) SearchRepository {
	return &searchRepository{db: db}
}

func (r *searchRepository) FindRoomBySearch(ctx context.Context, search *models.RoomSearch) ([]*models.Room, error) {
	if search == nil {
		return nil, errors.New("search is nil")
	}

	var rooms []*models.Room

	name := "%" + strings.ToLower(search.Name) + "%"
	desc := "%" + strings.ToLower(search.Description) + "%"
	category := "%" + strings.ToLower(search.Category.String()) + "%"

	rows := r.db.Where(`LOWER(name) LIKE ? AND LOWER(description) LIKE ? AND LOWER(category) LIKE ?`, name, desc, category).Find(&rooms)

	if rows.Error != nil {
		return nil, rows.Error
	}

	if rows.RowsAffected == 0 {
		return nil, errors.New(string(views.Err_NotFound))
	}

	return rooms, nil
}
