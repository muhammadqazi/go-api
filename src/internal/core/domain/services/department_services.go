package services

import (
  "github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/repositories"
)

type DepartmentServices interface {
	CreateDepartment(dtos.DepartmentCreateDTO) error
	FetchDepartment() error
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

func (s *departmentServices) CreateDepartment(department dtos.DepartmentCreateDTO) error {}
func (s *departmentServices) FetchDepartment() error{}

