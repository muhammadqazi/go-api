package services

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/repositories"
)

type FacultyServices interface {
	CreateFaculty(dtos.FacultyCreateDTO) error
	FetchFaculty(string) (entities.FacultiesEntity, error)
	FetchAllFaculties() ([]entities.FacultiesEntity, error)
}

type facultyServices struct {
	facultyMapper     mappers.FacultyMapper
	facultyRepository repositories.FacultyRepository
}

func NewFacultyServices(repo repositories.FacultyRepository, mapper mappers.FacultyMapper) FacultyServices {
	return &facultyServices{
		facultyRepository: repo,
		facultyMapper:     mapper,
	}
}

func (s *facultyServices) CreateFaculty(faculty dtos.FacultyCreateDTO) error {
	entity := s.facultyMapper.FacultyCreateMapper(faculty)
	if err := s.facultyRepository.InsertFaculty(entity); err != nil {
		return err
	}

	return nil

}
func (s *facultyServices) FetchFaculty(code string) (entities.FacultiesEntity, error) {
	return s.facultyRepository.SelectFacultyByCode(code)
}

func (s *facultyServices) FetchAllFaculties() ([]entities.FacultiesEntity, error) {
	return s.facultyRepository.QuertAllFaculties()
}
