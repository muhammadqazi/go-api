package entities

type FacultiesEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	FacultyID   uint   `gorm:"primaryKey;not null;uniqueIndex" json:"faculty_id"`
	Name        string `gorm:"type:varchar(50);not null" json:"name"`
	Code        string `gorm:"type:varchar(50);not null;uniqueIndex" json:"code"` // TODO
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	Dean        string `gorm:"type:varchar(50);not null" json:"dean"`
	ViceDean    string `gorm:"type:varchar(50);not null" json:"vice_dean"`
	Email       string `gorm:"type:varchar(50);not null" json:"email"`
	PhoneNumber string `gorm:"type:varchar(50);not null" json:"phone_number"`
	DeanEmail   string `gorm:"type:varchar(50);not null" json:"dean_email"`
	DeanPhone   string `gorm:"type:varchar(50);not null" json:"dean_phone"`

	DepartmentID uint `gorm:"not null" json:"department_id"`
}
