package usecase

import (
	"context"
	"zoomer/internal/search/presenter"
	"zoomer/internal/search/views"
)

type UseCase interface {
	SearchRooms(ctx context.Context, req *presenter.RoomSearchParams) *views.Response
}
