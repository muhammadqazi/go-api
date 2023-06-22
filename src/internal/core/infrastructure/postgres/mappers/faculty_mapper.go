package mappers

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
)

type FacultyMapper interface {
	FacultyCreateMapper(dto dtos.FacultyCreateDTO) entities.FacultiesEntity
}

type facultyMapper struct {
}

func NewFacultyMapper() FacultyMapper {
	return &facultyMapper{}
}

func (m *facultyMapper) FacultyCreateMapper(dto dtos.FacultyCreateDTO) entities.FacultiesEntity {
	return entities.FacultiesEntity{
		Name:        dto.Name,
		Description: dto.Description,
		Email:       dto.Email,
		PhoneNumber: dto.PhoneNumber,
	}
}
