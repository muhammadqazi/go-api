package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type StudentRepository interface {
	InsertStudent(student entities.StudentsEntity) (uint, error)
	QueryLastStudentID() (uint, error)
	QueryStudentByEmail(string) (entities.StudentsEntity, error)
	QueryStudentByID(uint) (entities.StudentsEntity, error)
	InsertStudentEnrollment(entities.StudentEnrollmentsEntity, uint) error
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

	res := r.conn.Create(&student)

	if res.Error != nil {
		return 0, res.Error
	}

	return student.StudentID, nil
}

func (r *studentConnection) QueryStudentByEmail(email string) (entities.StudentsEntity, error) {

	var student entities.StudentsEntity
	err := r.conn.Unscoped().Where("email = ?", email).First(&student).Error
	return student, err

}

func (r *studentConnection) QueryStudentByID(sid uint) (entities.StudentsEntity, error) {

	var student entities.StudentsEntity
	err := r.conn.Unscoped().Where("student_id = ?", sid).First(&student).Error
	return student, err

}
func (r *studentConnection) QueryLastStudentID() (uint, error) {

	var lastStudent entities.StudentsEntity

	rec := r.conn.Order("student_id DESC").Last(&lastStudent)

	if rec.Error != nil {
		return 0, rec.Error
	}

	return lastStudent.StudentID, nil

}

func (r *studentConnection) InsertStudentEnrollment(enrollment entities.StudentEnrollmentsEntity, supervisorID uint) error {

	tx := r.conn.Begin()

	if err := r.conn.Create(&enrollment).Error; err != nil {
		tx.Rollback()
		return err
	}

	var pivot entities.StudentCourseRequestEntity
	pivot.StudentEnrollmentID = enrollment.StudentEnrollmentID
	pivot.InstructorID = supervisorID

	if err := r.conn.Create(&pivot).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
