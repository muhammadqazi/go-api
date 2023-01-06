package entities

type FacultiesEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	FacultyID uint   `gorm:"primaryKey;not null;uniqueIndex" json:"faculty_id"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"`

	DepartmentsEntity []DepartmentsEntity `gorm:"foreignkey:faculty_id"`
	StudentsEntity    []StudentsEntity    `gorm:"foreignkey:faculty_id"`
}
