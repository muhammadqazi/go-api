package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type CurriculumRepository interface {
	InsertCurriculum(entities.CurriculumEntity, uint) error
}

type curriculumConnection struct {
	conn *gorm.DB
}

func NewCurriculumRepository(db *gorm.DB) CurriculumRepository {
	return &curriculumConnection{
		conn: db,
	}
}

func (r *curriculumConnection) InsertCurriculum(curriculum entities.CurriculumEntity, courseID uint) error {
	tx := r.conn.Begin()

	if err := tx.Create(&curriculum).Error; err != nil {
		tx.Rollback()
		return err
	}

	var pivot entities.CourseCurriculumEntity

	pivot.CurriculumID = curriculum.CurriculumID
	pivot.CourseID = courseID
	pivot.CreatedAt = curriculum.CreatedAt
	pivot.IsActive = curriculum.IsActive

	if err := tx.Create(&pivot).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
