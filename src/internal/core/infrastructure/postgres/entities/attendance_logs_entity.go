package entities

type AttendanceLogsEntity struct {
	StudentEnrollmentID uint `gorm:"primaryKey;not null;uniqueIndex" json:"student_enrollment_id"`
	IsAttended          bool `gorm:"not null" json:"is_attended"`
	AttendanceDate      uint `gorm:"not null" json:"attendance_date"`
	AttendanceTime      uint `gorm:"not null" json:"attendance_time"`

	StudentID uint `gorm:"not null" json:"student_id"`
	CourseID  uint `gorm:"not null" json:"course_id"`
}
