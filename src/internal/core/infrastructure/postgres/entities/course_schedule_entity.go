package entities

type CourseScheduleEntity struct {
	BaseEntity

	CourseScheduleID uint   `gorm:"primaryKey;not null;uniqueIndex" json:"course_schedule_id"`
	Day              string `gorm:"type:varchar(255);not null" json:"day"`
	StartTime        string `gorm:"type:varchar(255);not null" json:"start_time"`
	EndTime          string `gorm:"type:varchar(255);not null" json:"end_time"`
	LectureVenue     string `gorm:"type:varchar(255);not null" json:"room"`

	CourseID uint `gorm:"not null" json:"course_id"`
}
