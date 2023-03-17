package usecase

import (
	"zoomer/internal/search/presenter"
	"zoomer/internal/search/repository"
	"zoomer/internal/search/views"
)

type roomUsecase struct {
	roomRepo repository.RoomRepository
	searchRepo repository.SearchRepository
}

func NewSearchUseCase(searchRepo searchRepo, roomRepo repository.RoomRepository) UseCase {
	return &roomUsecase{
		searchRepo: searchRepo,
		roomRepo: roomRepo,
	}
}

func (r roomUsecase) SearchRooms(req *params.RoomSearch) *views.Response {
	search := req.ParseToModel()

	rooms, err := r.repo.FindRoombySearch(search)
	if err != nil {
		if err.Error() == string(views.Error_NotFound) {
			return views.Notfound(err)
		}
		return views.RepositoryError(err)
	}
	roomsView := views.NewRoomFind(rooms)
	return views.SuccessFindAll(roomsView)
}
