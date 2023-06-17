package entities

type BuildingsEntity struct {
	BaseEntity

	BuildingID    uint   `gorm:"primaryKey;not null;uniqueIndex" json:"building_id"`
	Name          string `gorm:"type:varchar(50);not null;uniqueIndex" json:"name"`
	Code          string `gorm:"type:varchar(50);not null;uniqueIndex" json:"code"`
	Description   string `gorm:"type:varchar(255);not null" json:"description"`
	NumberOfRooms int    `gorm:"not null" json:"number_of_rooms"`

	BuildingRoomsEntity []BuildingRoomsEntity `gorm:"foreignkey:BuildingID"`
}
