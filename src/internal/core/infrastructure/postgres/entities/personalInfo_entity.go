package entities

type PersonalInfoEntity struct {
	PersonalInfoID uint `gorm:"primaryKey;not null;uniqueIndex" json:"personal_info_id"`
	IDCardNumber   uint `gorm:"type:varchar(50);not null;uniqueIndex" json:"id_card_number"`
	PassportNumber uint `gorm:"type:varchar(50);not null;uniqueIndex" json:"passport_number"`
	FatherName     uint `gorm:"type:varchar(50);not null;uniqueIndex" json:"father_name"`
	MotherName     uint `gorm:"type:varchar(50);not null;uniqueIndex" json:"mother_name"`

	StudentID uint `gorm:"not null" json:"student_id"`
}
