package dtos

import "time"

type StudentCreateDTO struct {
	FirstName    string `json:"first_name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	Nationality  string `json:"nationality"`
	DOB          string `json:"dob"`
	PlaceOfBirth string `json:"place_of_birth"`
	Sex          string `json:"sex"`
	Password     string `json:"password"`
	Role         string `json:"role"`
	FacultyID    uint   `json:"faculty_id"`
	IsActive     bool   `json:"is_active"`
}

type StudentResponseDTO struct {
	StudentID           uint      `json:"student_id"`
	FirstName           string    `json:"first_name"`
	Surname             string    `json:"surname"`
	Email               string    `json:"email"`
	Nationality         string    `json:"nationality"`
	DOB                 string    `json:"dob"`
	PlaceOfBirth        string    `json:"place_of_birth"`
	Sex                 string    `json:"sex"`
	Role                string    `json:"role"`
	Status              string    `json:"status"`
	Semester            string    `json:"semester"`
	EnrollmentDate      time.Time `json:"enrollment_date"`
	GraduationDate      time.Time `json:"graduation_date"`
	FacultyID           uint      `json:"faculty_id"`
	PersonalInfoID      uint      `json:"personal_info_id"`
	ContactInfoID       uint      `json:"contact_info_id"`
	AddressID           uint      `json:"address_id"`
	StdAccountingInfoID uint      `json:"std_accounting_info_id"`
	IsActive            bool      `json:"is_active"`
	CreatedAt           time.Time `json:"created_at"`
	IsDeleted           bool      `json:"is_deleted"`
	IsGraduated         bool      `json:"is_graduated"`
}
