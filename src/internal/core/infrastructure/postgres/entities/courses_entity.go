package entities

type CoursesEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	CourseID    uint   `gorm:"primary_key;not null;uniqueIndex" json:"course_id"`
	CourseCode  string `gorm:"type:varchar(255);not null;uniqueIndex" json:"course_code"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	CreditHours int    `gorm:"type:int;not null" json:"credit_hours"`
	Quota       int    `gorm:"type:int;not null" json:"quota"`
	Status      string `gorm:"type:varchar(255);not null" json:"status"`
}
