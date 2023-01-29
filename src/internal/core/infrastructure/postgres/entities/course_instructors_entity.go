package entities

type CourseInstructorsEntity struct {
	BaseEntity

	CourseInstructorsID uint `gorm:"primaryKey;not null;uniqueIndex"`
	CourseID            uint `gorm:"not null"`
	InstructorID        uint `gorm:"not null"`
}
