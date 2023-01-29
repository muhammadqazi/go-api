package entities

type CurriculumEntity struct {
	BaseEntity

	CurriculumID uint   `gorm:"primaryKey;not null;uniqueIndex" json:"curriculum_id"`
	Semester     string `gorm:"not null" json:"semester"`
	Year         int    `gorm:"not null" json:"year"`

	DepartmentID uint `gorm:"not null" json:"department_id"`
	CourseID     uint `gorm:"not null" json:"course_id"`
}
