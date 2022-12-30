package entities

type AdvisorsEntity struct {
	BaseEntity // createdat, updatedat, deletedat, isactive

	AdvisorID     uint          `gorm:"primary_key;not null;uniqueIndex" json:"advisor_id"`
	Name          string        `gorm:"type:varchar(255);not null" json:"name"`
	Surname       string        `gorm:"type:varchar(255);not null" json:"surname"`
	Email         string        `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Password      string        `gorm:"type:varchar(255);not null" json:"password"`
	Office        string        `gorm:"type:varchar(255);not null" json:"office"`
	Line          string        `gorm:"type:varchar(255);not null" json:"line"`
	CourseIDs     []uint        `gorm:"type:uint" json:"course_id"`
	CoursesEntity CoursesEntity `gorm:"foreignKey:CourseID;AssociationForeignKey:CourseID"`
}
