package dtos

import "time"

type ExamScheduleCreateDTO struct {
	Date      time.Time `json:"date" validate:"required"`
	ExamType  string    `json:"exam_type" validate:"required"`
	Duration  int       `json:"duration" validate:"required"`
	ExamVenue string    `json:"exam_venue" validate:"required"`

	CourseID uint `json:"course_id" validate:"required"`
}

type ExamScheduleSchema struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	IsActive  bool      `gorm:"column:is_active"`
	ExamVenue string    `gorm:"column:exam_venue"`
	Date      time.Time `gorm:"column:date"`
	ExamType  string    `gorm:"column:exam_type"`
	Duration  int       `gorm:"column:duration"`
	CourseID  uint      `gorm:"column:course_id"`
	Code      string    `gorm:"column:code"`
	Name      string    `gorm:"column:name"`
	Credits   int       `gorm:"column:credits"`
}

type ExamScheduleFetchDTO struct {
	Schedule []ExamScheduleSchema
}
