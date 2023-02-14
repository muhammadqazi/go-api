package entities

type InstructorEnrollmentsEntity struct {
	BaseEntity

	InstructorEnrollmentID uint `gorm:"primaryKey;not null;uniqueIndex" json:"instructor_enrollment_id"`

	InstructorID uint `gorm:"not null" json:"instructor_id"`

	InstructorCoursesEntity []InstructorCoursesEntity `gorm:"foreignkey:InstructorEnrollmentID"`
}
