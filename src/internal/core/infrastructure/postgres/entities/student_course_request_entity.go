package entities

type StudentCourseRequestEntity struct {
	StudentCourseRequestID uint `gorm:"primaryKey;not null;uniqueIndex" json:"student_instructor_request_id"`

	InstructorID        uint `gorm:"not null" json:"instructor_id"`
	StudentEnrollmentID uint `gorm:"not null" json:"student_enrollment_id"`
}
