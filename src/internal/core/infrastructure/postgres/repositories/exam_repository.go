package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type ExamRepository interface {
	InsertExamSchedule(entity entities.ExamScheduleEntity) error
	QueryExam()
}

type examConnection struct {
	conn *gorm.DB
}

func NewExamRepository(db *gorm.DB) ExamRepository {
	return &examConnection{
		conn: db,
	}
}

func (r *examConnection) InsertExamSchedule(exam entities.ExamScheduleEntity) error {
	return r.conn.Create(&exam).Error
}
func (r *examConnection) QueryExam() {}
