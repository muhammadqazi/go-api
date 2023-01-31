package dtos

type curriculum struct {
	Semester string `json:"semester" binding:"required"`
	Year     int    `json:"year" binding:"required"`
	CourseID uint   `json:"course_id" binding:"required"`
}

type CurriculumCreateDTO struct {
	DepartmentID uint         `json:"department_id" validate:"required"`
	Curriculum   []curriculum `json:"curriculum" validate:"required,dive,required"`
}

type CurriculumSchema struct {
	Semester     string `json:"semester" binding:"required"`
	Year         int    `json:"year" binding:"required"`
	DepartmentID uint   `json:"department_id" binding:"required"`
}
