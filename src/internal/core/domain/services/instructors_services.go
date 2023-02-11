package services

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/repositories"
)

type InstructorsServices interface {
	CreateInstructors(dtos.InstructorCreateDTO) error
	FetchInstructorByEmail(string) (entities.InstructorsEntity, error)
	FetchInstructorByPhone(string) (entities.InstructorsEntity, error)
	FetchTermEnrollmentRequests(uint) ([]dtos.InstructorTermRequests, error)
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

func (s *instructorsServices) FetchInstructorByEmail(email string) (entities.InstructorsEntity, error) {
	return s.instructorsRepository.QueryInstructorByEmail(email)
}

func (s *instructorsServices) FetchInstructorByPhone(phone string) (entities.InstructorsEntity, error) {
	return s.instructorsRepository.QueryInstructorByPhone(phone)
}

func (s *instructorsServices) FetchTermEnrollmentRequests(id uint) ([]dtos.InstructorTermRequests, error) {
	return s.instructorsRepository.QueryTermEnrollmentRequests(id)
}
