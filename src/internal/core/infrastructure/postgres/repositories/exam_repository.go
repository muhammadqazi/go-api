package repositories

import (
	"errors"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"gorm.io/gorm"
)

type ExamRepository interface {
	InsertExamSchedule(entities.ExamScheduleEntity) error
	UpdateExamResults(entities.ExamResultsEntity, dtos.ExamResultsPatchDTO) error
}

type examConnection struct {
	conn       *gorm.DB
	examMapper mappers.ExamMapper
}

func NewExamRepository(db *gorm.DB, examMapper mappers.ExamMapper) ExamRepository {
	return &examConnection{
		conn:       db,
		examMapper: examMapper,
	}
}

func (r *examConnection) InsertExamSchedule(exam entities.ExamScheduleEntity) error {
	return r.conn.Create(&exam).Error
}

func (r *examConnection) UpdateExamResults(exam entities.ExamResultsEntity, dto dtos.ExamResultsPatchDTO) error {

	tx := r.conn.Begin()

	/*
		Fetch the existing result from the database, if it exists then update it, otherwise create a new one
	*/

	existingResult := entities.ExamResultsEntity{}
	var err error
	if err = tx.Where("student_id = ? AND semester = ? AND year = ?", exam.StudentID, exam.Semester, exam.Year).First(&existingResult).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := tx.Create(&exam).Error; err != nil {
			tx.Rollback()
			return err
		}
		existingResult = exam
	}

	courseResults := r.examMapper.ExamCourseResultsMapper(dto, existingResult.ExamResultID)
	if err := tx.Create(&courseResults).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}
