package entities

type StudentEnrollmentsEntity struct {
	BaseEntity

	StudentEnrollmentID uint   `gorm:"primaryKey;not null;uniqueIndex" json:"student_enrollment_id"`
	Semester            string `gorm:"not null" json:"semester"`
	Year                int    `gorm:"not null" json:"year"`
	IsApproved          bool   `gorm:"not null" json:"is_approved"`

	StudentID uint `gorm:"not null" json:"student_id"`
	CourseID  uint `gorm:"not null" json:"course_id"`

	StudentCourseRequestEntity []StudentCourseRequestEntity `gorm:"foreignkey:StudentEnrollmentID"`
}
