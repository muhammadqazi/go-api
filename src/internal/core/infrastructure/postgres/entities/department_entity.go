package entities

type DepartmentsEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	DepartmentID   uint   `gorm:"primaryKey;not null;uniqueIndex" json:"department_id"`
	DepartmentCode string `gorm:"type:varchar(255);not null;uniqueIndex" json:"department_code"`
	Name           string `gorm:"type:varchar(255);not null" json:"name"`
	FacultyID      uint   `gorm:"not null" json:"faculty_id"`

	CoursesEntity []CoursesEntity `gorm:"foreignkey:DepartmentID"`
}
