package services

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/repositories"
)

type StudentServices interface {
	CreateStudent()
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

func (s *studentServices) CreateStudent() {

}
