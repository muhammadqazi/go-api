package mappers

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"time"
)

type CurriculumMapper interface {
	CurriculumCreateMapper(schema dtos.CurriculumSchema) entities.CurriculumEntity
}

type curriculumMapper struct {
}

func NewCurriculumMapper() CurriculumMapper {
	return &curriculumMapper{}
}

func (s *curriculumMapper) CurriculumCreateMapper(curriculum dtos.CurriculumSchema) entities.CurriculumEntity {
	return entities.CurriculumEntity{
		Semester:     curriculum.Semester,
		Year:         curriculum.Year,
		DepartmentID: curriculum.DepartmentID,
		BaseEntity: entities.BaseEntity{
			CreatedAt: time.Now().UTC(),
			IsActive:  true,
		},
	}
}
