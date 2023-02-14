package entities

type StudentAttendanceEntity struct {
	BaseEntity

	StudentAttendanceID uint `gorm:"primaryKey;not null;uniqueIndex" json:"student_attendance_id"`

	Semester string `gorm:"type:varchar(30);not null" json:"semester"`
	Year     int    `gorm:"not null" json:"year"`

	StudentID uint `gorm:"not null" json:"student_id"`

	CourseAttendanceEntity []CourseAttendanceEntity `gorm:"foreignKey:StudentAttendanceID" json:"course_attendance_entity"`
}
