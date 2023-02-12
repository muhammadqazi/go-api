package dtos

import "time"

type CurriculumCreateDTO struct {
	DepartmentID uint         `json:"department_id" validate:"required"`
	Curriculum   []Curriculum `json:"curriculum" validate:"required,dive,required"`
}

type Curriculum struct {
	Semester   string `json:"semester" validate:"required"`
	Year       int    `json:"year" validate:"required"`
	CourseIDs  []uint `json:"course_ids" validate:"required,dive,required"`
	CourseLoad int    `json:"course_load" validate:"required"`
}

/* Using CurriculumSchema thing in the services */

type CurriculumSchema struct {
	Semester     string `json:"semester" validate:"required"`
	Year         int    `json:"year" validate:"required"`
	DepartmentID uint   `json:"department_id" validate:"required"`
}

type CourseCurriculumSchema struct {
	CourseID   uint `json:"course_id" validate:"required"`
	CourseLoad int  `json:"course_load" validate:"required"`
}

type CurriculumQueryReturnSchema struct {
	CourseID       uint      `json:"course_id"`
	CourseLoad     int       `json:"course_load"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
	CurriculumID   uint      `json:"curriculum_id"`
	Year           int       `json:"year"`
	Semester       string    `json:"semester"`
	DepartmentID   uint      `json:"department_id"`
	DepartmentName string    `json:"department_name"`
	DepartmentCode string    `json:"department_code"`
	NumberOfYears  int       `json:"number_of_years"`
	Code           string    `json:"code"`
	Name           string    `json:"name"`
	Credits        int       `json:"credits"`
	Ects           int       `json:"ects"`
	Practical      int       `json:"practical"`
	Theoretical    int       `json:"theoretical"`
}

/* These DTOs are for request from frontend select with department_id */

type CourseInfo struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Credits     int       `json:"credits"`
	Ects        int       `json:"ects"`
	Practical   int       `json:"practical"`
	Theoretical int       `json:"theoretical"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type CurriculumInfo struct {
	Semester   string `json:"semester"`
	Year       int    `json:"year"`
	CourseLoad int    `json:"course_load"`
	Courses    []CourseInfo
}

type CurriculumFetchDTO struct {
	DepartmentID   uint             `json:"department_id"`
	DepartmentName string           `json:"department_name"`
	DepartmentCode string           `json:"department_code"`
	NumberOfYears  int              `json:"number_of_years"`
	Curriculum     []CurriculumInfo `json:"curriculum"`
}
