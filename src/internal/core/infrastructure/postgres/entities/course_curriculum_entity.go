package entities

type CourseCurriculumEntity struct {
	BaseEntity

	CourseCurriculumID uint `gorm:"primaryKey;not null;uniqueIndex" json:"course_curriculum_id"`

	CourseID   uint   `gorm:"not null" json:"course_id"`
	CourseLoad int    `gorm:"not null" json:"course_load"`
	Semester   string `gorm:"not null" json:"semester"`
	Year       int    `gorm:"not null" json:"year"`

	CurriculumID uint `gorm:"not null" json:"curriculum_id"`
}
