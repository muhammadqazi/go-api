package dtos

import "time"

type StudentCreateDTO struct {
	FirstName      string `json:"first_name" validate:"required,min=2,max=50"`
	Surname        string `json:"surname" validate:"required,min=2,max=50"`
	Email          string `json:"email" validate:"required,email"`
	Nationality    string `json:"nationality" validate:"required,min=2,max=50"`
	DOB            string `json:"dob" validate:"required"`
	PlaceOfBirth   string `json:"place_of_birth" validate:"required,min=2,max=50"`
	Sex            string `json:"sex" validate:"required,min=2,max=50"`
	Password       string `json:"password" validate:"required,min=8,max=50"`
	Role           string `json:"role" validate:"required,min=2,max=50"`
	Scholarship    int    `json:"scholarship" validate:"required"`
	Discount       int    `json:"discount" validate:"required"`
	DiscountType   string `json:"discount_type" validate:"required,min=2,max=50"`
	AcceptanceType string `json:"acceptance_type" validate:"required,min=2,max=50"`
	DepartmentID   uint   `json:"department_id" validate:"required"`
	SupervisorID   uint   `json:"supervisor_id" validate:"required"`
}

type StudentSignInDTO struct {
	StudentID uint   `json:"student_id" validate:"required"`
	Password  string `json:"password" validate:"required"`
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

type TermRegistrationDTO struct {
	Semester string `json:"semester" validate:"required"`
	Year     int    `json:"year" validate:"required"`

	CourseIDs []uint `json:"course_ids" validate:"required,dive,required"`
}
