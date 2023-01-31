package entities

type CourseCurriculumEntity struct {
	BaseEntity

	CourseCurriculumID uint `gorm:"primaryKey;not null;uniqueIndex" json:"course_curriculum_id"`

	CourseID     uint `gorm:"not null" json:"course_id"`
	CurriculumID uint `gorm:"not null" json:"curriculum_id"`
}
