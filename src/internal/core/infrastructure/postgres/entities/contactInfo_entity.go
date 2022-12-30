package entities

type ContactInfoEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	ContactInfoID  uint   `gorm:"primary_key;not null;uniqueIndex" json:"contact_info_id"`
	Email          string `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	PhoneNumber    string `gorm:"type:varchar(255);not null;uniqueIndex" json:"phone_number"`
	LocalAddress   string `gorm:"type:varchar(255);not null;uniqueIndex" json:"address"`
	EmergencyName  string `gorm:"type:varchar(255);not null;uniqueIndex" json:"emergency_name"`
	EmergencyPhone string `gorm:"type:varchar(255);not null;uniqueIndex" json:"emergency_phone"`
}
