package entities

import "time"

type ExamScheduleEntity struct {
	BaseEntity

	ExamScheduleID uint `gorm:"primaryKey;not null;uniqueIndex" json:"exam_schedule_id"`

	Day       string    `gorm:"type:varchar(255);not null" json:"day"`
	StartTime time.Time `gorm:"type:timestamp;not null" json:"start_time"`
	EndTime   time.Time `gorm:"type:timestamp;not null" json:"end_time"`
	ExamVenue string    `gorm:"type:varchar(255);not null" json:"room"`

	CourseID uint `gorm:"not null" json:"course_id"`
}
