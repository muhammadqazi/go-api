package entities

type StudentEnrollmentsEntity struct {
	BaseEntity

	StudentEnrollmentID uint   `gorm:"primaryKey;not null;uniqueIndex" json:"student_enrollment_id"`
	Semester            string `gorm:"not null" json:"semester"`
	Year                int    `gorm:"not null" json:"year"`
	IsEnrolled          bool   `gorm:"not null;default:false" json:"is_enrolled"`

	StudentID    uint `gorm:"not null" json:"student_id"`
	InstructorID uint `gorm:"not null" json:"instructor_id"`

	StudentCourseRequestEntity []StudentCourseRequestEntity `gorm:"foreignkey:StudentEnrollmentID"`
}
