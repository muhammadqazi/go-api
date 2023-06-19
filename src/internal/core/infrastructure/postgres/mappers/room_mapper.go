package mappers

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
)

type RoomMapper interface {
	RoomCreateMapper(dto dtos.RoomCreateDTO) entities.BuildingRoomsEntity
}

type roomMapper struct {
}

func NewRoomMapper() RoomMapper {
	return &roomMapper{}
}

func (m *roomMapper) RoomCreateMapper(room dtos.RoomCreateDTO) entities.BuildingRoomsEntity {
	return entities.BuildingRoomsEntity{
		RoomNumber:      room.RoomNumber,
		RoomDescription: room.RoomDescription,
		NumberOfSeats:   room.NumberOfSeats,
		BuildingID:      room.BuildingID,
		IsAvailable:     room.IsAvailable,
	}
}
