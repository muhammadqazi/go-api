package repositories

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"gorm.io/gorm"
)

type FacultyRepository interface {
	InsertFaculty(entity entities.FacultiesEntity) error
	SelectFacultyByCode(code string) (entities.FacultiesEntity, error)
}

type facultyConnection struct {
	conn   *gorm.DB
	mapper mappers.FacultyMapper
}

func NewFacultyRepository(db *gorm.DB, mapper mappers.FacultyMapper) FacultyRepository {
	return &facultyConnection{
		conn:   db,
		mapper: mapper,
	}
}

func (r *facultyConnection) InsertFaculty(faculty entities.FacultiesEntity) error {
	if err := r.conn.Create(&faculty).Error; err != nil {
		return err
	}

	return nil
}
func (r *facultyConnection) SelectFacultyByCode(code string) (entities.FacultiesEntity, error) {
	var faculty entities.FacultiesEntity

	if err := r.conn.Where("code = ?", code).First(&faculty).Error; err != nil {
		return faculty, err
	}

	return faculty, nil
}
