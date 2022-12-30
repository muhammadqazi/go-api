package entities

type AddressesEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	AddressID uint   `gorm:"primary_key;not null;uniqueIndex" json:"address_id"`
	State     string `gorm:"type:varchar(255);not null" json:"state"`
	City      string `gorm:"type:varchar(255);not null" json:"city"`
	Province  string `gorm:"type:varchar(255)" json:"province"`
	Address   string `gorm:"type:varchar(255)" json:"address"`
}
