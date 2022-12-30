package entities

type DepartmentsEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	DepartmentID   uint          `gorm:"primary_key;not null;uniqueIndex" json:"department_id"`
	DepartmentCode string        `gorm:"type:varchar(255);not null;uniqueIndex" json:"department_code"`
	Name           string        `gorm:"type:varchar(255);not null" json:"name"`
	CourseIDs      []uint        `gorm:"type:uint" json:"course_id"`
	CoursesEntity  CoursesEntity `gorm:"foreignKey:CourseID;AssociationForeignKey:CourseID"`
}
