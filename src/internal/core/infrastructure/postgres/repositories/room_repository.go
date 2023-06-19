package repositories

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"gorm.io/gorm"
)

type RoomRepository interface {
	InsertRoom(entity entities.BuildingRoomsEntity) error
	QueryRoomByNumber(string) (entities.BuildingRoomsEntity, error)
	QueryAvailableRooms() ([]entities.BuildingRoomsEntity, error)
}

type roomConnection struct {
	conn   *gorm.DB
	mapper mappers.RoomMapper
}

func NewRoomRepository(db *gorm.DB, mapper mappers.RoomMapper) RoomRepository {
	return &roomConnection{
		conn:   db,
		mapper: mapper,
	}
}

func (r *roomConnection) InsertRoom(room entities.BuildingRoomsEntity) error {
	return r.conn.Create(&room).Error
}
func (r *roomConnection) QueryRoomByNumber(number string) (entities.BuildingRoomsEntity, error) {
	var room entities.BuildingRoomsEntity
	err := r.conn.Where("room_number = ?", number).First(&room).Error
	return room, err
}

func (r *roomConnection) QueryAvailableRooms() ([]entities.BuildingRoomsEntity, error) {
	var rooms []entities.BuildingRoomsEntity
	err := r.conn.Where("is_available = ?", true).Find(&rooms).Error
	return rooms, err
}
