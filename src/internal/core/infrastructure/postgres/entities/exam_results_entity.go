package entities

type JSONB []interface{}

type ExamResultsEntity struct {
	BaseEntity

	ExamResultsID uint  `gorm:"primaryKey;not null;uniqueIndex" json:"exam_results_id"`
	Results       JSONB `gorm:"type:jsonb;default:'[]'" json:"results"`

	CourseID  uint `gorm:"not null" json:"course_id"`
	StudentID uint `gorm:"not null" json:"student_id"`
}
