package repository

import (
	"context"
	"zoomer/internal/models"
)

type ChatRepository interface {
	UpdateContactList(ctx context.Context, username, contact string) error

	CreateChat(ctx context.Context, c *models.Chat) (string, error)
}
