package dtos

type CurriculumCreateDTO struct {
	DepartmentID uint         `json:"department_id" validate:"required"`
	Curriculum   []curriculum `json:"curriculum" validate:"required,dive,required"`
}

type curriculum struct {
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
