package entities

type InstructorEnrollmentsEntity struct {
	BaseEntity

	InstructorEnrollmentID uint `gorm:"primaryKey;not null;uniqueIndex" json:"instructor_enrollment_id"`

	Semester string `gorm:"type:varchar(255);not null" json:"semester"`
	Year     string `gorm:"type:varchar(255);not null" json:"year"`

	CourseID     uint `gorm:"not null" json:"course_id"`
	InstructorID uint `gorm:"not null" json:"instructor_id"`
}

