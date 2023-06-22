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
	FirstName    string  `json:"first_name" validate:"required,min=3,max=255"`
	LastName     string  `json:"last_name" validate:"required,min=3,max=255"`
	PhoneNumber  string  `json:"phone_number" validate:"required,min=3,max=255"`
	Email        string  `json:"email" validate:"required,email"`
	Password     string  `json:"password"`
	DOB          string  `json:"dob" validate:"required,min=3,max=255"`
	PlaceOfBirth string  `json:"place_of_birth" validate:"required,min=3,max=255"`
	Sex          string  `json:"sex" validate:"required"`
	Nationality  string  `json:"nationality" validate:"required`
	Role         role    `json:"role" validate:"required"`
	Salary       float64 `json:"salary" validate:"required"`
	OfficeId     uint    `json:"office_id" validate:"required"`
}

type InstructorSignInDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type InstructorTermRequests struct {
	EnrollmentID      uint      `json:"enrollment_id"`
	SupervisorName    string    `json:"supervisor_name"`
	SupervisorSurname string    `json:"supervisor_surname"`
	SupervisorID      uint      `json:"supervisor_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at"`
	IsEnrolled        bool      `json:"is_enrolled"`
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
}

type InstructorTermRequestsFetchDTO struct {
	EnrollmentID      uint                 `json:"enrollment_id"`
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
	Courses           []CourseApprovalInfo `json:"courses"`
}

type InstructorApproveEnrollmentRequestDTO struct {
	EnrollmentID uint  `json:"enrollment_id" validate:"required"`
	IsDeclined   *bool `json:"is_declined" validate:"required"`
}

/* Enroll Instructor to courses */

type InstructorCourseEnrollmentDTO struct {
	InstructorID uint   `json:"instructor_id"`
	Semester     string `json:"semester"`
	Year         int    `json:"year"`
	CourseIDs    []uint `json:"course_ids"`
}

/* Course Enrolments for instructors schema */

type InstructorEnrollmentsSchema struct {
	EnrollmentDate   time.Time `gorm:"column:enrollment_date"`
	IsActive         bool      `gorm:"column:is_active"`
	InstructorID     uint      `gorm:"column:instructor_id"`
	FirstName        string    `gorm:"column:first_name"`
	LastName         string    `gorm:"column:last_name"`
	InstructorEmail  string    `gorm:"column:email"`
	InstructorStatus bool      `gorm:"column:instructor_status"`
	CourseID         uint      `gorm:"column:course_id"`
	Day              string    `gorm:"column:day"`
	StartTime        string    `gorm:"column:start_time"`
	EndTime          string    `gorm:"column:end_time"`
	LectureVenue     string    `gorm:"column:lecture_venue"`
	CourseScheduleID uint      `gorm:"column:course_schedule_id"`
	CourseName       string    `gorm:"column:name"`
	CourseCode       string    `gorm:"column:code"`
	Credits          int       `gorm:"column:credits"`
	Theoretical      int       `gorm:"column:theoretical"`
	Practical        int       `gorm:"column:practical"`
}

type CourseEnrollmentInfo struct {
	EnrollmentDate   time.Time `json:"enrollment_date"`
	EnrollmentStatus bool      `json:"enrollment_status"`
	ID               uint      `json:"id"`
	Name             string    `json:"name"`
	Code             string    `json:"code"`
	Credits          int       `json:"credits"`
	Practical        int       `json:"practical"`
	Theoretical      int       `json:"theoretical"`
	Day              string    `json:"day"`
	StartTime        string    `json:"start_time"`
	EndTime          string    `json:"end_time"`
	LectureVenue     string    `json:"lecture_venue"`
	CourseScheduleID uint      `json:"course_schedule_id"`
}

type LecturesEnrollmentInfo struct {
	Day     string                 `json:"day"`
	Courses []CourseEnrollmentInfo `json:"courses"`
}

type InstructorEnrollmentsFetchDTO struct {
	InstructorID     uint                     `json:"instructor_id"`
	FirstName        string                   `json:"first_name"`
	LastName         string                   `json:"last_name"`
	InstructorEmail  string                   `json:"email"`
	InstructorStatus bool                     `json:"instructor_status"`
	Lectures         []LecturesEnrollmentInfo `json:"lectures"`
}

type SupervisedStudentSchema struct {
	StudentID             uint   `json:"student_id"`
	FirstName             string `json:"first_name"`
	Surname               string `json:"surname"`
	Email                 string `json:"email"`
	Nationality           string `json:"nationality"`
	DOB                   string `json:"dob"`
	Sex                   string `json:"sex"`
	Role                  string `json:"role"`
	Status                string `json:"status"`
	AccessStatus          string `json:"access_status"`
	DepartmentID          uint   `json:"department_id"`
	DepartmentCode        string `json:"department_code"`
	DepartmentName        string `json:"department_name"`
	NumberOfYears         uint   `json:"number_of_years"`
	DepartmentDescription string `json:"department_description"`
	FacultyID             uint   `json:"faculty_id"`
	FacultyName           string `json:"faculty_name"`
}

type SupervisedStudentsDTO struct {
	StudentID    uint   `json:"student_id"`
	FirstName    string `json:"first_name"`
	SurName      string `json:"surname"`
	Email        string `json:"email"`
	Nationality  string `json:"nationality"`
	DOB          string `json:"dob"`
	Sex          string `json:"sex"`
	Role         string `json:"role"`
	Status       string `json:"status"`
	AccessStatus string `json:"access_status"`
}

type RegisteredStudentsDTO struct {
	Year           int    `json:"year"`
	Semester       string `json:"semester"`
	IsEnrolled     bool   `json:"is_enrolled"`
	ApprovedAt     string `json:"approved_at"`
	DeclinedAt     string `json:"declined_at"`
	Name           string `json:"name"`
	StudentID      uint   `json:"student_id"`
	DepartmentCode string `json:"department_code"`
	DepartmentName string `json:"department_name"`
}
