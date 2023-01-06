package entities

import "time"

type StudentsEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	StudentID               uint      `gorm:"primary_key;not null;uniqueIndex" json:"student_id"`
	FirstName               string    `gorm:"type:varchar(255);not null" json:"first_name"`
	Surname                 string    `gorm:"type:varchar(255);not null" json:"surname"`
	Email                   string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Nationality             string    `gorm:"type:varchar(255);not null" json:"nationality"`
	DOB                     string    `gorm:"type:varchar(255);not null" json:"dob"`
	PlaceOfBirth            string    `gorm:"type:varchar(255);not null" json:"place_of_birth"`
	Sex                     string    `gorm:"type:varchar(255);not null" json:"sex"`
	Password                string    `gorm:"type:varchar(255);not null" json:"password"`
	Role                    string    `gorm:"type:varchar(255);not null" json:"role"`
	Status                  string    `gorm:"type:varchar(255);not null" json:"status"`
	Semester                string    `gorm:"type:varchar(255);not null" json:"semester"`
	EnrollmentDate          time.Time `gorm:"type:timestamp" json:"enrollment_date"`
	GraduationDate          time.Time `gorm:"type:timestamp" json:"graduation_date"`
	FacultyID               uint
	FacultyEntity           FacultiesEntity `gorm:"foreignKey:FacultyID;AssociationForeignKey:FacultyID"`
	PersonalInfoID          uint
	PersonalInfoEntity      PersonalInfoEntity `gorm:"foreignKey:PersonalInfoID;AssociationForeignKey:PersonalInfoID"`
	ContactInfoID           uint
	ContactInfoEntity       ContactInfoEntity `gorm:"foreignKey:ContactInfoID;AssociationForeignKey:ContactInfoID"`
	AddressID               uint
	AddressEntity           AddressesEntity `gorm:"foreignKey:AddressID;AssociationForeignKey:AddressID"`
	StdAccountingInfoID     uint
	StdAccountingInfoEntity StdAccountingInfoEntity `gorm:"foreignKey:StdAccountingInfoID;AssociationForeignKey:StdAccountingInfoID"`
	IsGraduated             bool                    `gorm:"type:boolean;not null" json:"is_graduated"`
	IsDeleted               bool                    `gorm:"type:boolean;not null" json:"is_deleted"`
}
