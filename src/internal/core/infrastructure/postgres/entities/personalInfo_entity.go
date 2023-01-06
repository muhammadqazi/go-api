package entities

type PersonalInfoEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	PersonalInfoID uint `gorm:"primaryKey;not null;uniqueIndex" json:"personal_info_id"`
	IDCardNumber   uint `gorm:"type:varchar(255);not null;uniqueIndex" json:"id_card_number"`
	PassportNumber uint `gorm:"type:varchar(255);not null;uniqueIndex" json:"passport_number"`
	FatherName     uint `gorm:"type:varchar(255);not null;uniqueIndex" json:"father_name"`
	MotherName     uint `gorm:"type:varchar(255);not null;uniqueIndex" json:"mother_name"`

	StudentsEntity []StudentsEntity `gorm:"foreignkey:PersonalInfoID"`
}
