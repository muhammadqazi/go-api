package entities

type ContactInfoEntity struct {
	BaseEntity

	ContactInfoID  uint   `gorm:"primaryKey;not null;uniqueIndex" json:"contact_info_id"`
	Email          string `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	PhoneNumber    string `gorm:"type:varchar(255);not null;uniqueIndex" json:"phone_number"`
	LocalAddress   string `gorm:"type:varchar(255);not null;uniqueIndex" json:"address"`
	EmergencyName  string `gorm:"type:varchar(255);not null;uniqueIndex" json:"emergency_name"`
	EmergencyPhone string `gorm:"type:varchar(255);not null;uniqueIndex" json:"emergency_phone"`

	StudentID uint `gorm:"not null" json:"student_id"`
}
