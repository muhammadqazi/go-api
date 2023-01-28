package entities

type DepartmentsEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	DepartmentID   uint   `gorm:"primaryKey;not null;uniqueIndex" json:"department_id"`
	DepartmentCode string `gorm:"type:varchar(50);not null;uniqueIndex" json:"department_code"`
	Name           string `gorm:"type:varchar(50);not null" json:"name"`
	Description    string `gorm:"type:varchar(255);not null" json:"description"`
	Dean           string `gorm:"type:varchar(50);not null" json:"dean"`
	ViceDean       string `gorm:"type:varchar(50);not null" json:"vice_dean"`
	Email          string `gorm:"type:varchar(50);not null" json:"email"`
	PhoneNumber    string `gorm:"type:varchar(50);not null" json:"phone_number"`
	DeanEmail      string `gorm:"type:varchar(50);not null" json:"dean_email"`
	DeanPhone      string `gorm:"type:varchar(50);not null" json:"dean_phone"`

	HeadID uint `gorm:"column:head_id" json:"head_id"`

	StudentsEntity   []StudentsEntity   `gorm:"foreignkey:DepartmentID"`
	FacultiesEntity  []FacultiesEntity  `gorm:"foreignkey:DepartmentID"`
	CurriculumEntity []CurriculumEntity `gorm:"foreignkey:DepartmentID"`
}
