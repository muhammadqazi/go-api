package services

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/repositories"
)

type DepartmentServices interface {
	CreateDepartment(dtos.DepartmentCreateDTO) error
	FetchDepartmentByCode(string) (entities.DepartmentsEntity, error)
	FetchAllDepartments() ([]entities.DepartmentsEntity, error)
}

type departmentServices struct {
	departmentMapper     mappers.DepartmentMapper
	departmentRepository repositories.DepartmentRepository
}

func NewDepartmentServices(repo repositories.DepartmentRepository, mapper mappers.DepartmentMapper) DepartmentServices {
	return &departmentServices{
		departmentRepository: repo,
		departmentMapper:     mapper,
	}
}

func (s *departmentServices) CreateDepartment(department dtos.DepartmentCreateDTO) error {
	departmentEntity := s.departmentMapper.DepartmentCreateMapper(department)
	return s.departmentRepository.InsertDepartment(departmentEntity)
}
func (s *departmentServices) FetchDepartmentByCode(code string) (entities.DepartmentsEntity, error) {
	return s.departmentRepository.QueryDepartmentById(code)
}

func (s *departmentServices) FetchAllDepartments() ([]entities.DepartmentsEntity, error) {
	return s.departmentRepository.QueryAllDepartments()
}
