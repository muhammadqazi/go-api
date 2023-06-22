package entities

type BuildingRoomsEntity struct {
	BaseEntity

	BuildingRoomID  uint   `gorm:"primaryKey;not null;uniqueIndex" json:"building_room_id"`
	RoomNumber      string `gorm:"type:varchar(50);not null;uniqueIndex" json:"room_number"`
	RoomDescription string `gorm:"type:varchar(255);not null" json:"room_description"`
	NumberOfSeats   uint   `gorm:"not null" json:"number_of_seats"`
	IsAvailable     bool   `gorm:"not null" json:"is_available"`

	BuildingID uint `gorm:"not null" json:"building_id"`

	InstructorsEntity []InstructorsEntity `gorm:"foreignkey:OfficeId"`
}
