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
