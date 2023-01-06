package entities

type CoursesEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	CourseID    uint   `gorm:"primaryKey;not null;uniqueIndex" json:"course_id"`
	CourseCode  string `gorm:"type:varchar(255);not null;uniqueIndex" json:"course_code"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	CreditHours int    `gorm:"type:int;not null" json:"credit_hours"`
	Quota       int    `gorm:"type:int;not null" json:"quota"`
	Status      string `gorm:"type:varchar(255);not null" json:"status"`

	DepartmentID uint `gorm:"not null" json:"department_id"`

	AdvisorsEntity []AdvisorsEntity `gorm:"foreignkey:CourseID"`
}
