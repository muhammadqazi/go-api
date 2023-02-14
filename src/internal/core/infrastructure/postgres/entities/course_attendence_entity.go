package entities

import "time"

type CourseAttendanceEntity struct {
	CourseAttendanceID uint `gorm:"primaryKey;not null;uniqueIndex" json:"course_attendance_id"`
	
	LectureTime time.Time `gorm:"type:timestamp;not null" json:"lecture_time"`
	IsAttended  bool      `gorm:"not null" json:"is_attended"`

	StudentAttendanceID uint `gorm:"not null" json:"student_attendance_id"`
	CourseID            uint `gorm:"not null" json:"course_id"`
}
