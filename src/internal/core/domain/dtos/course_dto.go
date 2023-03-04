package dtos

type CourseSchedule struct {
	Day          string `json:"day" validate:"required"`
	StartTime    string `json:"start_time" validate:"required"`
	EndTime      string `json:"end_time" validate:"required"`
	LectureVenue string `json:"lecture_venue" validate:"required"`
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
	Name         string `gorm:"column:name"`
	Code         string `gorm:"column:code"`
	Description  string `gorm:"column:description"`
	Credits      int    `gorm:"column:credits"`
	Ects         int    `gorm:"column:ects"`
	Theoretical  int    `gorm:"column:theoretical"`
	Practical    int    `gorm:"column:practical"`
	IsActive     bool   `gorm:"column:is_active"`
	Day          string `gorm:"column:day"`
	StartTime    string `gorm:"column:start_time"`
	EndTime      string `gorm:"column:end_time"`
	LectureVenue string `gorm:"column:lecture_venue"`
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
