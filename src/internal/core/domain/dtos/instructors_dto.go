package dtos

import "time"

type role string

const (
	RootRole       role = "root"
	AdminRole      role = "admin"
	InstructorRole role = "instructor"
	AssistantRole  role = "assistant"
)

type InstructorCreateDTO struct {
	FirstName    string `json:"first_name" validate:"required,min=3,max=255"`
	LastName     string `json:"last_name" validate:"required,min=3,max=255"`
	PhoneNumber  string `json:"phone_number" validate:"required,min=3,max=255"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8,max=255"`
	DOB          string `json:"dob" validate:"required,min=3,max=255"`
	PlaceOfBirth string `json:"place_of_birth" validate:"required,min=3,max=255"`
	Sex          string `json:"sex" validate:"required"`
	Nationality  string `json:"nationality" validate:"required`
	Role         role   `json:"role" validate:"required"`
}

type InstructorSignInDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type InstructorTermRequests struct {
	RequestID         uint      `json:"request_id"`
	SupervisorName    string    `json:"supervisor_name"`
	SupervisorSurname string    `json:"supervisor_surname"`
	SupervisorID      uint      `json:"supervisor_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at"`
	IsApproved        bool      `json:"is_approved"`
	Semester          string    `json:"semester"`
	Year              int       `json:"year"`
	StudentID         uint      `json:"student_id"`
	StudentName       string    `json:"student_name"`
	StudentSurname    string    `json:"student_surname"`
	StudentStatus     string    `json:"student_status"`
	AccessStatus      string    `json:"access_status"`
	CourseID          uint      `json:"course_id"`
	CourseName        string    `json:"course_name"`
	CourseCode        string    `json:"course_code"`
	CourseCredits     int       `json:"course_credits"`
	CourseStatus      bool      `json:"course_status"`
	ECTS              int       `json:"ects"`
	Theoretical       int       `json:"theoretical"`
	Practical         int       `json:"practical"`
}

type CourseApprovalInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Credits     int    `json:"credits"`
	Ects        int    `json:"ects"`
	Practical   int    `json:"practical"`
	Theoretical int    `json:"theoretical"`
	IsApproved  bool   `json:"is_approved"`
	RequestID   uint   `json:"request_id"`
}

type InstructorTermRequestsFetchDTO struct {
	SupervisorID      uint                 `json:"supervisor_id"`
	SupervisorName    string               `json:"supervisor_name"`
	SupervisorSurname string               `json:"supervisor_surname"`
	StudentID         uint                 `json:"student_id"`
	StudentName       string               `json:"student_name"`
	StudentSurname    string               `json:"student_surname"`
	StudentStatus     string               `json:"student_status"`
	AccessStatus      string               `json:"access_status"`
	Semester          string               `json:"semester"`
	Year              int                  `json:"year"`
	IsApproved        bool                 `json:"is_approved"`
	Courses           []CourseApprovalInfo `json:"courses"`
}

type InstructorApproveEnrollmentRequestDTO struct {
	StudentID   uint   `json:"student_id" validate:"required"`
	RequestsIDs []uint `json:"request_ids" validate:"required,dive"`
}
