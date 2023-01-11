package dtos

import "time"

type StudentCreateDTO struct {
	FirstName      string `json:"first_name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	Nationality    string `json:"nationality"`
	DOB            string `json:"dob"`
	PlaceOfBirth   string `json:"place_of_birth"`
	Sex            string `json:"sex"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	IsActive       bool   `json:"is_active"`
	FacultyID      uint   `json:"faculty_id"`
	PersonalInfoID uint   `json:"personalinfo_id"`
	ContactInfoID  uint   `json:"contactinfo_id"`
	AddressID      uint   `json:"address_id"`
	AccountsID     uint   `json:"accounts_id"`
}

type StudentSignInDTO struct {
	StudentID uint   `json:"student_id"`
	Password  string `json:"password"`
}

type StudentResponseDTO struct {
	StudentID      uint      `json:"student_id"`
	FirstName      string    `json:"first_name"`
	Surname        string    `json:"surname"`
	Email          string    `json:"email"`
	Nationality    string    `json:"nationality"`
	DOB            string    `json:"dob"`
	PlaceOfBirth   string    `json:"place_of_birth"`
	Sex            string    `json:"sex"`
	Role           string    `json:"role"`
	Status         string    `json:"status"`
	Semester       string    `json:"semester"`
	EnrollmentDate time.Time `json:"enrollment_date"`
	GraduationDate time.Time `json:"graduation_date"`
	FacultyID      uint      `json:"faculty_id"`
	PersonalInfoID uint      `json:"personalinfo_id"`
	ContactInfoID  uint      `json:"contactinfo_id"`
	AddressID      uint      `json:"address_id"`
	AccountsID     uint      `json:"accounts_id"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	IsDeleted      bool      `json:"is_deleted"`
	IsGraduated    bool      `json:"is_graduated"`
}
