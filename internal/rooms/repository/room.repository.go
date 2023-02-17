package repository

import (
	"context"
	"gorm.io/gorm"
	"zoomer/internal/models"
)

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (cr *roomRepository) CreateRoom(ctx context.Context, room *models.Room) error {
	result := cr.db.WithContext(ctx).Create(&room)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *roomRepository) GetRoomByUserId(ctx context.Context, userId string) ([]*models.Room, error) {
	var rooms []*models.Room
	err := cr.db.WithContext(ctx).Where(&models.Room{
		CreatedBy: userId}).Find(&rooms).Error

	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (cr *roomRepository) GetAllRooms(ctx context.Context) ([]*models.Room, error) {
	var chats []*models.Room
	err := cr.db.WithContext(ctx).Limit(200).Find(&chats).Error

	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (cr *roomRepository) CountRooms(ctx context.Context, userId string) (int, error) {
	var count int

	err := cr.db.WithContext(ctx).Raw(`SELECT COUNT(*) FROM chats WHERE chats.created_by = ? AND DATE_TRUNC('day', "created_at") =  CURRENT_DATE GROUP BY DATE_TRUNC('day', "created_at")`, userId).Scan(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
