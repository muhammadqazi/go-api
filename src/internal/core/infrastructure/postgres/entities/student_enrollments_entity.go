package entities

type StudentEnrollmentsEntity struct {
	BaseEntity

	StudentEnrollmentID uint `gorm:"primaryKey;not null;uniqueIndex" json:"student_enrollment_id"`
	Semester            uint `gorm:"not null" json:"semester"`
	Year                uint `gorm:"not null" json:"year"`

	StudentID uint `gorm:"not null" json:"student_id"`
	CourseID  uint `gorm:"not null" json:"course_id"`
}
