package dtos

type CourseSchedule struct {
	Day           string `json:"day" validate:"required"`
	StartTime     string `json:"start_time" validate:"required"`
	EndTime       string `json:"end_time" validate:"required"`
	LectureVenue  string `json:"lecture_venue" validate:"required"`
	IsTheoretical *bool  `json:"is_theoretical" validate:"required"`
}

type CourseCreateDTO struct {
	Name           string           `json:"name" validate:"required"`
	Code           string           `json:"code" validate:"required"`
	Description    string           `json:"description" validate:"required"`
	Credits        int              `json:"credits" validate:"required"`
	ECTS           int              `json:"ects" validate:"required"`
	Theoretical    int              `json:"theoretical" validate:"required"`
	Practical      int              `json:"practical" validate:"required"`
	CourseSchedule []CourseSchedule `json:"course_schedule" validate:"required,dive,required"`
}

type CourseUpdateDTO struct {
	Name        string `json:"name" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Description string `json:"description" validate:"required"`
	Credits     int    `json:"credits" validate:"required"`
	ECTS        int    `json:"ects" validate:"required"`
	Theoretical int    `json:"theoretical" validate:"required"`
	Practical   int    `json:"practical" validate:"required"`
}

type CourseFetchByCodeSchema struct {
	Name          string `gorm:"column:name"`
	Code          string `gorm:"column:code"`
	Description   string `gorm:"column:description"`
	Credits       int    `gorm:"column:credits"`
	Ects          int    `gorm:"column:ects"`
	Theoretical   int    `gorm:"column:theoretical"`
	Practical     int    `gorm:"column:practical"`
	IsActive      bool   `gorm:"column:is_active"`
	Day           string `gorm:"column:day"`
	StartTime     string `gorm:"column:start_time"`
	EndTime       string `gorm:"column:end_time"`
	LectureVenue  string `gorm:"column:lecture_venue"`
	IsTheoretical bool   `gorm:"column:is_theoretical"`
}

type CourseFetchByCodeDTO struct {
	Name           string           `json:"name"`
	Code           string           `json:"code"`
	Description    string           `json:"description"`
	Credits        int              `json:"credits"`
	ECTS           int              `json:"ects"`
	Theoretical    int              `json:"theoretical"`
	Practical      int              `json:"practical"`
	IsActive       bool             `json:"is_active"`
	CourseSchedule []CourseSchedule `json:"course_schedule"`
}

type CourseLogSchema struct {
	StudentID        uint   `gorm:"column:student_id"`
	PracticalHours   int    `gorm:"practical_hours"`
	TheoreticalHours int    `gorm:"theoretical_hours"`
	Credits          int    `gorm:"credits"`
	CourseID         uint   `gorm:"course_id"`
	Day              string `gorm:"day"`
	StartTime        string `gorm:"start_time"`
	EndTime          string `gorm:"end_time"`
	IsTheoretical    bool   `gorm:"is_theoretical"`
}

type CourseInstructorUpdateDTO struct {
	CourseId      uint   `json:"course_id" validate:"required"`
	InstructorIds []uint `json:"instructor_ids" validate:"required min=1"`
}
