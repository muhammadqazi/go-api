package entities

type JSONB interface{}

type ExamCourseResultsEntity struct {
	ExamCourseResultsID uint  `gorm:"primaryKey;not null;uniqueIndex" json:"exam_course_results_id"`
	Results             JSONB `gorm:"type:jsonb;default:'{}'" json:"results"`

	ExamResultID uint `gorm:"not null" json:"exam_result_id"`
	CourseID     uint `gorm:"not null" json:"course_id"`
}
