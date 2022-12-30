package entities

type FacultiesEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	FacultyID         uint              `gorm:"primary_key;not null;uniqueIndex" json:"faculty_id"`
	Name              string            `gorm:"type:varchar(255);not null" json:"name"`
	DepartmentIDs     []uint            `gorm:"type:uint" json:"department_id"`
	DepartmentsEntity DepartmentsEntity `gorm:"foreignKey:DepartmentID;AssociationForeignKey:DepartmentID"`
}
