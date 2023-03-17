package usecase

type UseCase interface {
	SearchRooms(req *presenter.RoomSearch) *views.Response
}
