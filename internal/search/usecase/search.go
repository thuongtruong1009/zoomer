package usecase

import (
	"context"
	"zoomer/internal/search/presenter"
	"zoomer/internal/search/repository"
	"zoomer/internal/search/views"

	room "zoomer/internal/rooms/repository"
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
