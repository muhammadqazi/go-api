package entities

type CoursesEntity struct {
	BaseEntity

	CourseID    uint   `gorm:"primaryKey;not null;uniqueIndex" json:"course_id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Code        string `gorm:"type:varchar(255);not null" json:"code"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	Credits     int    `gorm:"not null" json:"credits"`
	ECTS        int    `gorm:"not null" json:"ects"`
	Theoretical int    `gorm:"not null" json:"theoretical"`
	Practical   int    `gorm:"not null" json:"practical"`

	CourseInstructorsEntity  []CourseInstructorsEntity  `gorm:"foreignkey:CourseID"`
	CurriculumEntity         []CurriculumEntity         `gorm:"foreignkey:CourseID"`
	TranscriptEntity         []TranscriptEntity         `gorm:"foreignkey:CourseID"`
	CourseScheduleEntity     []CourseScheduleEntity     `gorm:"foreignkey:CourseID"`
	ExamScheduleEntity       []ExamScheduleEntity       `gorm:"foreignkey:CourseID"`
	ExamResultsEntity        []ExamResultsEntity        `gorm:"foreignkey:CourseID"`
	StudentEnrollmentsEntity []StudentEnrollmentsEntity `gorm:"foreignkey:CourseID"`
	AttendanceLogsEntity     []AttendanceLogsEntity     `gorm:"foreignkey:CourseID"`
}
