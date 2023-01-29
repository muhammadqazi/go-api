package services

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/repositories"
)

type CurriculumServices interface {
	CreateCurriculum()
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

func (s *curriculumServices) CreateCurriculum() {

}
