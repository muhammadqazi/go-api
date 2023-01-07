package entities

type AddressesEntity struct {
	AddressID uint   `gorm:"primaryKey;not null;uniqueIndex" json:"address_id"`
	State     string `gorm:"type:varchar(255);not null" json:"state"`
	City      string `gorm:"type:varchar(255);not null" json:"city"`
	Province  string `gorm:"type:varchar(255)" json:"province"`
	Address   string `gorm:"type:varchar(255)" json:"address"`

	StudentsEntity []StudentsEntity `gorm:"foreignkey:AddressID"`
}
