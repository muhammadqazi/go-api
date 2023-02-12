package entities

type CurriculumEntity struct {
	BaseEntity

	CurriculumID uint `gorm:"primaryKey;not null;uniqueIndex" json:"curriculum_id"`

	DepartmentID uint `gorm:"not null" json:"department_id"`

	CourseCurriculumEntity []CourseCurriculumEntity `gorm:"foreignKey:CurriculumID"`
}
