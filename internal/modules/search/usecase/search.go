package usecase

import (
	"context"
	"github.com/thuongtruong1009/zoomer/internal/modules/search/presenter"
	"github.com/thuongtruong1009/zoomer/internal/modules/search/repository"
	"github.com/thuongtruong1009/zoomer/internal/modules/search/views"

	room "github.com/thuongtruong1009/zoomer/internal/modules/rooms/repository"
)

type roomUsecase struct {
	searchRepo repository.SearchRepository
	roomRepo   room.RoomRepository
}

func NewSearchUseCase(searchRepo repository.SearchRepository, roomRepo room.RoomRepository) UseCase {
	return &roomUsecase{
		searchRepo: searchRepo,
		roomRepo:   roomRepo,
	}
}

func (r *roomUsecase) SearchRooms(ctx context.Context, req *presenter.RoomSearchParams) *views.Response {
	search := req.ParseToModel()

	rooms, err := r.searchRepo.FindRoomBySearch(ctx, search)
	if err != nil {
		if err.Error() == string(views.Err_NotFound) {
			return views.NotFound(err)
		}
		return views.RepositoryError(err)
	}
	roomsView := views.NewRoomsFind(ctx, rooms)
	return views.SuccessFindAll(roomsView)
}
