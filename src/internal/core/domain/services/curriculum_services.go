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
	/*
		"""
		curriculum.Curriculum is an array of [course ids, semester and year], in the curriculum entity
		we will store the semester, year and department id, and in the curriculum course entity which is a
		pivot table we will store the course id and curriculum id.

		CurriculumSchema is the blueprint of the curriculum entity
		"""
	*/
	for _, v := range curriculum.Curriculum {
		schema := dtos.CurriculumSchema{
			Semester:     v.Semester,
			Year:         v.Year,
			DepartmentID: curriculum.DepartmentID,
		}
		m := s.curriculumMapper.CurriculumCreateMapper(schema)

		for _, id := range v.CourseIDs {
			pivot := dtos.CourseCurriculumSchema{
				CourseID:   id,
				CourseLoad: v.CourseLoad,
			}

			if err := s.curriculumRepository.InsertCurriculum(m, pivot); err != nil {
				return err
			}
		}

	}
	return nil
}
