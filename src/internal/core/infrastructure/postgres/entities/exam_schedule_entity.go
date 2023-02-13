package entities

import "time"

type ExamScheduleEntity struct {
	BaseEntity

	ExamScheduleID uint `gorm:"primaryKey;not null;uniqueIndex" json:"exam_schedule_id"`

	Date      time.Time `gorm:"type:timestamp;not null" json:"date"`
	ExamType  string    `gorm:"type:varchar(255);not null" json:"exam_type"`
	Duration  int       `gorm:"not null" json:"duration"`
	ExamVenue string    `gorm:"type:varchar(255);not null" json:"room"`

	CourseID uint `gorm:"not null" json:"course_id"`
}
