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
	Installments   int    `json:"installments"`
	MotherName     string `json:"mother_name" validate:"required,min=2,max=50"`
	FatherName     string `json:"father_name" validate:"required,min=2,max=50"`
	IDCardNumber   string `json:"id_card_number" validate:"min=2,max=50"`
	PassportNumber string `json:"passport_number" validate:"min=2,max=50"`
}

type StudentPatchDTO struct {
	FirstName      string `json:"first_name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	Nationality    string `json:"nationality"`
	DOB            string `json:"dob"`
	PlaceOfBirth   string `json:"place_of_birth"`
	Sex            string `json:"sex"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	Scholarship    int    `json:"scholarship"`
	Discount       int    `json:"discount"`
	DiscountType   string `json:"discount_type"`
	AcceptanceType string `json:"acceptance_type"`
	DepartmentID   uint   `json:"department_id"`
	SupervisorID   uint   `json:"supervisor_id"`
	Installments   int    `json:"installments"`
	MotherName     string `json:"mother_name"`
	FatherName     string `json:"father_name"`
	IDCardNumber   string `json:"id_card_number"`
	PassportNumber string `json:"passport_number"`
}

type StudentsFetchDTO struct {
	FirstName      string `json:"first_name"`
	Surname        string `json:"surname"`
	Email          string `json:"email"`
	Nationality    string `json:"nationality"`
	DOB            string `json:"dob"`
	PlaceOfBirth   string `json:"place_of_birth"`
	Sex            string `json:"sex"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	Scholarship    int    `json:"scholarship"`
	Discount       int    `json:"discount"`
	DiscountType   string `json:"discount_type"`
	AcceptanceType string `json:"acceptance_type"`
	DepartmentID   uint   `json:"department_id"`
	SupervisorID   uint   `json:"supervisor_id"`
	Installments   int    `json:"installments"`
	MotherName     string `json:"mother_name"`
	FatherName     string `json:"father_name"`
	IDCardNumber   string `json:"id_card_number"`
	PassportNumber string `json:"passport_number"`
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
	CourseIDs []uint `json:"course_ids" validate:"required,dive,required"`
}

/* Student Timetable DTOs */

type TimetableSchema struct {
	StudentEnrollmentID uint   `gorm:"column:student_enrollment_id"`
	StudentID           uint   `gorm:"column:student_id"`
	Year                int    `gorm:"column:year"`
	Semester            string `gorm:"column:semester"`
	CourseID            uint   `gorm:"column:course_id"`
	RequestID           uint   `gorm:"column:student_course_request_id"`
	Name                string `gorm:"column:name"`
	Code                string `gorm:"column:code"`
	Day                 string `gorm:"column:day"`
	StartTime           string `gorm:"column:start_time"`
	EndTime             string `gorm:"column:end_time"`
	Credits             int    `gorm:"column:credits"`
	LectureVenue        string `gorm:"column:lecture_venue"`
	IsTheoretical       bool   `gorm:"column:is_theoretical"`
}

type LectureInfo struct {
	CourseID      uint   `json:"course_id"`
	CourseCode    string `json:"course_code"`
	CourseName    string `json:"course_name"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	LectureVenue  string `json:"lecture_venue"`
	Credits       int    `json:"credits"`
	IsTheoretical bool   `json:"is_theoretical"`
}

type TimeTableInfo struct {
	Day      string        `json:"day"`
	Lectures []LectureInfo `json:"lectures"`
}

type TimetableFetchDTO struct {
	StudentID uint   `json:"student_id"`
	Year      int    `json:"year"`
	Semester  string `json:"semester"`
	Timetable []TimeTableInfo
}

/* Student Attendance DTOs */

type StudentAttendancePatchDTO struct {
	StudentID   uint      `json:"student_id" validate:"required"`
	CourseID    uint      `json:"course_id" validate:"required"`
	LectureTime time.Time `json:"lecture_time" validate:"required"`
	IsAttended  *bool     `json:"is_attended" validate:"required"`
}

/* Student Attendance Fetch */

type StudentAttendanceSchema struct {
	LectureTime   string `gorm:"column:lecture_time"`
	Day           string `gorm:"column:day"`
	StartTime     string `gorm:"column:start_time"`
	EndTime       string `gorm:"column:end_time"`
	IsAttended    bool   `gorm:"column:is_attended"`
	IsTheoretical bool   `gorm:"column:is_theoretical"`
	CourseCode    string `gorm:"column:course_code"`
	CourseName    string `gorm:"column:course_name"`
	Credits       int    `gorm:"column:credits"`
}

type StudentAttendanceFetchDTO struct {
	LectureTIme                          string  `json:"lecture_time"`
	Day                                  string  `json:"day"`
	StartTime                            string  `json:"start_time"`
	EndTime                              string  `json:"end_time"`
	IsAttended                           bool    `json:"is_attended"`
	IsTheoretical                        bool    `json:"is_theoretical"`
	CourseCode                           string  `json:"course_code"`
	CourseName                           string  `json:"course_name"`
	Credits                              int     `json:"credits"`
	TotalLectures                        int     `json:"total_lectures"`
	AttendedLectures                     int     `json:"attended_lectures"`
	TotalTheoreticalLectures             int     `json:"total_theoretical_lectures"`
	AttendedTheoreticalLectures          int     `json:"attended_theoretical_lectures"`
	TotalPracticalLectures               int     `json:"total_practical_lectures"`
	AttendedPracticalLectures            int     `json:"attended_practical_lectures"`
	TotalLectureAttendancePercentage     float64 `json:"total_lecture_attendance_percentage"`
	TotalTheoreticalAttendancePercentage float64 `json:"total_theoretical_attendance_percentage"`
	TotalPracticalAttendancePercentage   float64 `json:"total_practical_attendance_percentage"`
}

/* Reset Password */

type ResetPasswordDTO struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=20"`
}

/* Forgot Password */

type ForgotPasswordRequestDTO struct {
	StudentID uint `json:"student_id" validate:"required"`
}

type ForgotPasswordVerifyDTO struct {
	Code int `json:"code" validate:"required"`
}

type NewPasswordDTO struct {
	Password string `json:"password" validate:"required,min=8,max=20"`
}
