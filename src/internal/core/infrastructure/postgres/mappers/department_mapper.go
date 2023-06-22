package mappers

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
)

type DepartmentMapper interface {
	DepartmentCreateMapper(dto dtos.DepartmentCreateDTO) entities.DepartmentsEntity
}

type departmentMapper struct {
}

func NewDepartmentMapper() DepartmentMapper {
	return &departmentMapper{}
}

func (m *departmentMapper) DepartmentCreateMapper(department dtos.DepartmentCreateDTO) entities.DepartmentsEntity {
	return entities.DepartmentsEntity{
		Name:          department.Name,
		Description:   department.Description,
		Email:         department.Email,
		PhoneNumber:   department.PhoneNumber,
		NumberOfYears: department.NumberOfYears,
		FacultyID:     department.FacultyID,
		HeadID:        department.HeadID,
	}
}
