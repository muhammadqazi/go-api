package entities

type InstructorCoursesEntity struct {
	InstructorCourseID uint `gorm:"primaryKey;not null;uniqueIndex" json:"instructor_course_enrollment_id"`

	Semester string `gorm:"type:varchar(255);not null" json:"semester"`
	Year     int    `gorm:"not null" json:"year"`
	
	InstructorEnrollmentID uint `gorm:"not null" json:"instructor_enrollment_id"`
	CourseID               uint `gorm:"not null" json:"course_id"`
}
