package services

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/repositories"
)

type InstructorsServices interface {
	CreateInstructors(dtos.InstructorCreateDTO) error
}

type instructorsServices struct {
	instructorsMapper     mappers.InstructorsMapper
	instructorsRepository repositories.InstructorsRepository
}

func NewInstructorsServices(repo repositories.InstructorsRepository, mapper mappers.InstructorsMapper) InstructorsServices {
	return &instructorsServices{
		instructorsRepository: repo,
		instructorsMapper:     mapper,
	}
}

func (s *instructorsServices) CreateInstructors(instructor dtos.InstructorCreateDTO) error {

	m := s.instructorsMapper.InstructorCreateMapper(instructor)
	return s.instructorsRepository.InsertInstructors(m)
}
