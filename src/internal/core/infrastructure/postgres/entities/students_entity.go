package entities

import "time"

type StudentsEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	StudentID      uint      `gorm:"primaryKey;not null;uniqueIndex" json:"student_id"`
	FirstName      string    `gorm:"type:varchar(50);not null" json:"first_name"`
	Surname        string    `gorm:"type:varchar(50);not null" json:"surname"`
	Email          string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"email"`
	Nationality    string    `gorm:"type:varchar(50);not null" json:"nationality"`
	DOB            string    `gorm:"type:varchar(50);not null" json:"dob"`
	PlaceOfBirth   string    `gorm:"type:varchar(50);not null" json:"place_of_birth"`
	Sex            string    `gorm:"type:varchar(50);not null" json:"sex"`
	Password       string    `gorm:"type:varchar(255);not null" json:"password"`
	Role           string    `gorm:"type:varchar(50);not null" json:"role"`
	Status         string    `gorm:"type:varchar(50);not null" json:"status"`
	AccessStatus   string    `gorm:"type:varchar(50);not null" json:"access_status"`
	AcceptanceType string    `gorm:"type:varchar(50);not null" json:"acceptance_type"`
	Semester       string    `gorm:"type:varchar(50);not null" json:"semester"`
	GraduationDate time.Time `gorm:"type:timestamp" json:"graduation_date"`

	IsGraduated bool `gorm:"type:boolean;not null;default:false" json:"is_graduated"`

	//SupervisorID uint `gorm:"column:instructor_id" json:"supervisor_id"`
	DepartmentID uint `gorm:"column:department_id" json:"department_id"`

	AccountsEntity     []AccountsEntity     `gorm:"foreignkey:StudentID"`
	PersonalInfoEntity []PersonalInfoEntity `gorm:"foreignkey:StudentID"`
	ContactInfoEntity  []ContactInfoEntity  `gorm:"foreignkey:StudentID"`
	TranscriptEntity   []TranscriptEntity   `gorm:"foreignkey:StudentID"`
	ExamResultsEntity  []ExamResultsEntity  `gorm:"foreignkey:StudentID"`
}
