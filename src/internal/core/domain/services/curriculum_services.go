package services

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/repositories"
)

type CurriculumServices interface {
	CreateCurriculum(dto dtos.CurriculumCreateDTO) error
}

type curriculumServices struct {
	curriculumMapper     mappers.CurriculumMapper
	curriculumRepository repositories.CurriculumRepository
}

func NewCurriculumServices(repo repositories.CurriculumRepository, mapper mappers.CurriculumMapper) CurriculumServices {
	return &curriculumServices{
		curriculumRepository: repo,
		curriculumMapper:     mapper,
	}
}

func (s *curriculumServices) CreateCurriculum(curriculum dtos.CurriculumCreateDTO) error {

	for _, v := range curriculum.Curriculum {
		schema := dtos.CurriculumSchema{
			Semester:     v.Semester,
			Year:         v.Year,
			DepartmentID: curriculum.DepartmentID,
		}
		m := s.curriculumMapper.CreateCurriculum(schema)
		if err := s.curriculumRepository.InsertCurriculum(m, v.CourseID); err != nil {
			return err
		}
	}
	return nil
}
