package repositories

import "gorm.io/gorm"

type StudentRepository interface {
	InsertStudent()
}

type studentConnection struct {
	conn *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentConnection{
		conn: db,
	}
}

func (r *studentConnection) InsertStudent() {
}
