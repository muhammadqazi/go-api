package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type CurriculumRepository interface {
	InsertCurriculum(entities.CurriculumEntity, dtos.CourseCurriculumSchema) error
}

type curriculumConnection struct {
	conn *gorm.DB
}

func NewCurriculumRepository(db *gorm.DB) CurriculumRepository {
	return &curriculumConnection{
		conn: db,
	}
}

func (r *curriculumConnection) InsertCurriculum(curriculum entities.CurriculumEntity, courseCurriculum dtos.CourseCurriculumSchema) error {
	tx := r.conn.Begin()

	if err := tx.Create(&curriculum).Error; err != nil {
		tx.Rollback()
		return err
	}

	var pivot entities.CourseCurriculumEntity
	pivot.CurriculumID = curriculum.CurriculumID
	pivot.CourseLoad = courseCurriculum.CourseLoad
	pivot.CourseID = courseCurriculum.CourseID

	if err := tx.Create(&pivot).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
