package entities

import "time"

type CourseAttendanceEntity struct {
	CourseAttendanceID uint `gorm:"primaryKey;not null;uniqueIndex" json:"course_attendance_id"`

	LectureTime   time.Time `gorm:"type:timestamp;not null" json:"lecture_time"`
	Day           string    `gorm:"type:varchar(255);not null" json:"day"`
	StartTime     string    `gorm:"type:varchar(255);not null" json:"start_time"`
	EndTime       string    `gorm:"type:varchar(255);not null" json:"end_time"`
	IsAttended    bool      `gorm:"not null" json:"is_attended"`
	IsTheoretical bool      `gorm:"not null" json:"is_theoretical"`

	StudentAttendanceID uint `gorm:"not null" json:"student_attendance_id"`
	CourseID            uint `gorm:"not null" json:"course_id"`
}
