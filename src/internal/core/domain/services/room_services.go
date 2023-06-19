package services

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/repositories"
)

type RoomServices interface {
	CreateRoom(dtos.RoomCreateDTO) error
	FetchRoomByNumber(string) (entities.BuildingRoomsEntity, error)
	FetchAvailableRooms() ([]entities.BuildingRoomsEntity, error)
}

type roomServices struct {
	roomMapper     mappers.RoomMapper
	roomRepository repositories.RoomRepository
}

func NewRoomServices(repo repositories.RoomRepository, mapper mappers.RoomMapper) RoomServices {
	return &roomServices{
		roomRepository: repo,
		roomMapper:     mapper,
	}
}

func (s *roomServices) CreateRoom(room dtos.RoomCreateDTO) error {
	entity := s.roomMapper.RoomCreateMapper(room)
	return s.roomRepository.InsertRoom(entity)
}
func (s *roomServices) FetchRoomByNumber(number string) (entities.BuildingRoomsEntity, error) {
	return s.roomRepository.QueryRoomByNumber(number)
}

func (s *roomServices) FetchAvailableRooms() ([]entities.BuildingRoomsEntity, error) {
	return s.roomRepository.QueryAvailableRooms()
}
