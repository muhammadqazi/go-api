package entities

import "time"

type StudentsEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	StudentID      uint      `gorm:"primaryKey;not null;uniqueIndex" json:"student_id"`
	FirstName      string    `gorm:"type:varchar(255);not null" json:"first_name"`
	Surname        string    `gorm:"type:varchar(255);not null" json:"surname"`
	Email          string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Nationality    string    `gorm:"type:varchar(255);not null" json:"nationality"`
	DOB            string    `gorm:"type:varchar(255);not null" json:"dob"`
	PlaceOfBirth   string    `gorm:"type:varchar(255);not null" json:"place_of_birth"`
	Sex            string    `gorm:"type:varchar(255);not null" json:"sex"`
	Password       string    `gorm:"type:varchar(255);not null" json:"password"`
	Role           string    `gorm:"type:varchar(255);not null" json:"role"`
	Status         string    `gorm:"type:varchar(255);not null" json:"status"`
	Semester       string    `gorm:"type:varchar(255);not null" json:"semester"`
	EnrollmentDate time.Time `gorm:"type:timestamp" json:"enrollment_date"`
	GraduationDate time.Time `gorm:"type:timestamp" json:"graduation_date"`

	FacultyID      uint `gorm:"column:faculty_id;not null" json:"faculty_id"`
	PersonalInfoID uint `gorm:"null" json:"personalinfo_id"`
	ContactInfoID  uint `gorm:"null" json:"contactinfo_id"`
	AddressID      uint `gorm:"null" json:"address_id"`
	AccountsID     uint `gorm:"not null" json:"accounts_id"`

	IsGraduated bool `gorm:"type:boolean;not null" json:"is_graduated"`
	IsDeleted   bool `gorm:"type:boolean;not null" json:"is_deleted"`
}
