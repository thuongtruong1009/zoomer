package usecase

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/modules/search/presenter"
	"github.com/thuongtruong1009/zoomer/internal/modules/search/views"
)

type UseCase interface {
	SearchRooms(ctx context.Context, req *presenter.RoomSearchParams) *views.Response
}
