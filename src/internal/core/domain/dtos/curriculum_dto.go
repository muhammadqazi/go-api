package dtos

type curriculum struct {
	Semester string `json:"semester" validate:"required"`
	Year     int    `json:"year" validate:"required"`
	CourseID uint   `json:"course_id" validate:"required"`
}

type CurriculumCreateDTO struct {
	DepartmentID uint         `json:"department_id" validate:"required"`
	Curriculum   []curriculum `json:"curriculum" validate:"required,dive,required"`
}

/* Using CurriculumSchema thing in the services */

type CurriculumSchema struct {
	Semester     string `json:"semester" validate:"required"`
	Year         int    `json:"year" validate:"required"`
	DepartmentID uint   `json:"department_id" validate:"required"`
}
