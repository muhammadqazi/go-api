package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type StudentRepository interface {
	InsertStudent(student entities.StudentsEntity) (uint, error)
	GetLastStudentID() (uint, error)
	GetStudentByEmail(email string) (entities.StudentsEntity, error)
	GetStudentByStudentID(sid uint) (entities.StudentsEntity, error)
	AddFaculty(f entities.FacultiesEntity)
}

type studentConnection struct {
	conn *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentConnection{
		conn: db,
	}
}

func (r *studentConnection) InsertStudent(student entities.StudentsEntity) (uint, error) {

	rec := r.conn.Create(&student)

	if rec.Error != nil {
		return 0, rec.Error
	}

	return student.StudentID, nil
}

func (r *studentConnection) GetStudentByEmail(email string) (entities.StudentsEntity, error) {

	var student entities.StudentsEntity
	err := r.conn.Unscoped().Where("email = ?", email).First(&student).Error
	return student, err

}

func (r *studentConnection) GetStudentByStudentID(sid uint) (entities.StudentsEntity, error) {

	var student entities.StudentsEntity
	err := r.conn.Unscoped().Where("student_id = ?", sid).First(&student).Error
	return student, err

}

func (r *studentConnection) AddFaculty(f entities.FacultiesEntity) {

	r.conn.Create(&f)

}

func (r *studentConnection) GetLastStudentID() (uint, error) {

	var lastStudent entities.StudentsEntity

	rec := r.conn.Order("student_id DESC").Last(&lastStudent)

	if rec.Error != nil {
		return 0, rec.Error
	}

	return lastStudent.StudentID, nil

}
