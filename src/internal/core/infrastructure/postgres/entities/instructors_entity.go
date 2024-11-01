package entities

type InstructorsEntity struct {
	BaseEntity

	InstructorID uint   `gorm:"primaryKey;not null;uniqueIndex" json:"instructor_id"`
	FirstName    string `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName     string `gorm:"type:varchar(255);not null" json:"last_name"`
	Email        string `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	PhoneNumber  string `gorm:"type:varchar(255);not null;uniqueIndex" json:"phone_number"`
	Password     string `gorm:"type:varchar(255);not null" json:"password"`
	DOB          string `gorm:"type:varchar(255);not null" json:"dob"`
	PlaceOfBirth string `gorm:"type:varchar(255);not null" json:"place_of_birth"`
	Sex          string `gorm:"type:varchar(255);not null" json:"sex"`
	Nationality  string `gorm:"type:varchar(255);not null" json:"nationality"`
	Role         string `gorm:"type:varchar(255);not null" json:"role"`
	IsVerified   bool   `gorm:"type:boolean;not null;default:false" json:"is_verified"`
	// Office       string `gorm:"type:varchar(255);not null" json:"office"`
	// OfficeLine  string `gorm:"type:varchar(255);not null" json:"office_line"`
	// OfficeHours string `gorm:"type:varchar(255);not null" json:"office_hours"`
	Salary   float64 `gorm:"type:float;not null" json:"salary"`
	OfficeId uint    `gorm:"column:office_id" json:"office_id"`

	InstructorEnrollmentsEntity []InstructorEnrollmentsEntity `gorm:"foreignkey:InstructorID"`
	DepartmentsEntity           []DepartmentsEntity           `gorm:"foreignkey:HeadID"`
	FacultiesEntity             []FacultiesEntity             `gorm:"foreignkey:DeanID"`
	FacultiesEntityVice         []FacultiesEntity             `gorm:"foreignkey:ViceDeanId"`
	StudentsEntity              []StudentsEntity              `gorm:"foreignkey:SupervisorID"`

	StudentEnrollmentsEntity []StudentEnrollmentsEntity `gorm:"foreignkey:SupervisorID"`
}
