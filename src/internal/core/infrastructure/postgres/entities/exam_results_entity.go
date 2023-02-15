package entities

type ExamResultsEntity struct {
	BaseEntity

	ExamResultID uint `gorm:"primaryKey;not null;uniqueIndex" json:"exam_results_id"`

	Semester  string `gorm:"type:varchar(255);not null" json:"semester"`
	Year      int    `gorm:"not null" json:"year"`
	StudentID uint   `gorm:"not null" json:"student_id"`

	ExamCourseResultsEntity []ExamCourseResultsEntity `gorm:"foreignKey:ExamResultID" json:"exam_course_results"`
}
