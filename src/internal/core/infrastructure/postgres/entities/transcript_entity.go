package entities

type TranscriptEntity struct {
	BaseEntity

	TranscriptID uint    `gorm:"primaryKey;not null;uniqueIndex" json:"transcript_id"`
	Grade        string  `gorm:"type:varchar(255);not null" json:"grade"`
	Semester     int     `gorm:"not null" json:"semester"`
	Year         int     `gorm:"not null" json:"year"`
	Credit       int     `gorm:"not null" json:"credit"`
	CGPA         float32 `gorm:"not null" json:"cgpa"`

	CourseID  uint `gorm:"not null" json:"course_id"`
	StudentID uint `gorm:"not null" json:"student_id"`
}
