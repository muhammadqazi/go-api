package entities

type FacultiesEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	FacultyID uint   `gorm:"primaryKey;not null;uniqueIndex" json:"faculty_id"`
	Name      string `gorm:"type:varchar(50);not null" json:"name"`
	//Code        string `gorm:"type:varchar(50);not null;uniqueIndex" json:"code"` // TODO
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	//Dean        string `gorm:"type:varchar(50);not null" json:"dean"`
	//ViceDean    string `gorm:"type:varchar(50);not null" json:"vice_dean"`
	Email       string `gorm:"type:varchar(50);not null" json:"email"`
	PhoneNumber string `gorm:"type:varchar(50);not null" json:"phone_number"`
	//DeanEmail   string `gorm:"type:varchar(50);not null" json:"dean_email"`
	//DeanPhone   string `gorm:"type:varchar(50);not null" json:"dean_phone"`

	DeanID     uint `gorm:"column:dean_id" json:"dean_id"`
	ViceDeanId uint `gorm:"column:vice_dean_id" json:"vice_dean_id"`

	DepartmentsEntity []DepartmentsEntity `gorm:"foreignkey:FacultyID"`
}
