package services

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/repositories"
)

type StudentServices interface {
	CreateStudent(dtos.StudentCreateDTO, uint, string) (uint, error)
	FetchLastStudentID() (uint, error)
	FetchStudentByEmail(string) (entities.StudentsEntity, error)
	FetchStudentByID(uint) (entities.StudentsEntity, error)
}

type studentServices struct {
	studentMapper     mappers.StudentMapper
	studentRepository repositories.StudentRepository
}

func NewStudentServices(repo repositories.StudentRepository, mapper mappers.StudentMapper) StudentServices {
	return &studentServices{
		studentRepository: repo,
		studentMapper:     mapper,
	}
}

func (s *studentServices) CreateStudent(student dtos.StudentCreateDTO, sid uint, semester string) (uint, error) {
	m := s.studentMapper.StudentCreateMapper(student, sid, semester)
	return s.studentRepository.InsertStudent(m)
}

func (s *studentServices) FetchStudentByEmail(email string) (entities.StudentsEntity, error) {
	return s.studentRepository.QueryStudentByEmail(email)
}

func (s *studentServices) FetchStudentByID(sid uint) (entities.StudentsEntity, error) {
	return s.studentRepository.QueryStudentByID(sid)
}

func (s *studentServices) FetchLastStudentID() (uint, error) {
	return s.studentRepository.QueryLastStudentID()
}
