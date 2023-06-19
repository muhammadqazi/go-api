package entities

type InstructorEnrollmentsEntity struct {
	BaseEntity

	InstructorEnrollmentID uint `gorm:"primaryKey;not null;uniqueIndex" json:"instructor_enrollment_id"`

	InstructorID uint `gorm:"not null" json:"instructor_id"`

	Semester string `gorm:"type:varchar(255);not null" json:"semester"`
	Year     int    `gorm:"not null" json:"year"`

	InstructorCoursesEntity []InstructorCoursesEntity `gorm:"foreignkey:InstructorEnrollmentID"`
}
