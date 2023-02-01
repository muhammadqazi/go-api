package mappers

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"time"
)

type InstructorsMapper interface {
	InstructorCreateMapper(dtos.InstructorCreateDTO) entities.InstructorsEntity
}

type instructorsMapper struct {
}

func NewInstructorsMapper() InstructorsMapper {
	return &instructorsMapper{}
}

func (s *instructorsMapper) InstructorCreateMapper(instructor dtos.InstructorCreateDTO) entities.InstructorsEntity {
	return entities.InstructorsEntity{
		FirstName:    instructor.FirstName,
		LastName:     instructor.LastName,
		Email:        instructor.Email,
		Password:     instructor.Password,
		PhoneNumber:  instructor.PhoneNumber,
		DOB:          instructor.DOB,
		PlaceOfBirth: instructor.PlaceOfBirth,
		Sex:          instructor.Sex,
		Nationality:  instructor.Nationality,
		Role:         string(instructor.Role),
		BaseEntity: entities.BaseEntity{
			CreatedAt: time.Now().UTC(),
			IsActive:  true,
		},
	}
}
