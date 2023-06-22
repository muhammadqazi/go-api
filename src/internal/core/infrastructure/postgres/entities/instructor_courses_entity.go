package entities

type InstructorCoursesEntity struct {
	InstructorCourseID uint `gorm:"primaryKey;not null;uniqueIndex" json:"instructor_course_enrollment_id"`

	InstructorEnrollmentID uint `gorm:"not null" json:"instructor_enrollment_id"`
	CourseID               uint `gorm:"not null" json:"course_id"`
}
