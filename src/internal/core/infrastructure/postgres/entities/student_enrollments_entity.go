package entities

type StudentEnrollmentsEntity struct {
	BaseEntity

	StudentEnrollmentID uint   `gorm:"primaryKey;not null;uniqueIndex" json:"student_enrollment_id"`
	Semester            string `gorm:"not null" json:"semester"`
	Year                int    `gorm:"not null" json:"year"`
	IsEnrolled          bool   `gorm:"not null;default:false" json:"is_enrolled"`
	ApprovedAt          string `gorm:"not null" json:"approved_at"`
	DeclinedAt          string `gorm:"not null" json:"declined_at"`

	StudentID    uint `gorm:"not null" json:"student_id"`
	SupervisorID uint `gorm:"not null" json:"supervisor_id"`

	StudentCourseRequestEntity []StudentCourseRequestEntity `gorm:"foreignkey:StudentEnrollmentID"`
}
