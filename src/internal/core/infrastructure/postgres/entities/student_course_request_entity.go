package entities

import "time"

type StudentCourseRequestEntity struct {
	StudentCourseRequestID uint `gorm:"primaryKey;not null;uniqueIndex" json:"student_instructor_request_id"`

	IsApproved bool `gorm:"not null" json:"is_approved"`

	ApprovedAt time.Time `gorm:"type:timestamp;null" json:"approved_at"`
	DeclinedAt time.Time `gorm:"type:timestamp;null" json:"declined_at"`

	CourseID            uint `gorm:"not null" json:"course_id"`
	StudentEnrollmentID uint `gorm:"not null" json:"student_enrollment_id"`
}
