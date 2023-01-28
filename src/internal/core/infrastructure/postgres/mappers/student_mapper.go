package mappers

import (
	"time"

	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/security"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
)

type StudentMapper interface {
	StudentCreateMapper(dtos.StudentCreateDTO, uint, string) entities.StudentsEntity
	StudentResponseMapper(entities.StudentsEntity) dtos.StudentResponseDTO
}

type studentMapper struct {
}

func NewStudentMapper() StudentMapper {
	return &studentMapper{}
}

const (
	New               = "New"
	Active            = "Active"
	InActive          = "InActive"
	Pending           = "Pending"
	FullAccess        = "FullAccess"
	ProvisionalAccess = "ProvisionalAccess"
	NoAccess          = "NoAccess"
)

const (
	Role = "Student"
)

func (m *studentMapper) StudentCreateMapper(student dtos.StudentCreateDTO, sid uint, semester string) entities.StudentsEntity {

	hashedPassword, _ := security.HashPassword(student.Password)

	return entities.StudentsEntity{
		StudentID:    sid,
		FirstName:    student.FirstName,
		Surname:      student.Surname,
		Email:        student.Email,
		Nationality:  student.Nationality,
		DOB:          student.DOB,
		PlaceOfBirth: student.PlaceOfBirth,
		Sex:          student.Sex,
		Password:     hashedPassword,
		Role:         Role,
		Semester:     semester,
		BaseEntity: entities.BaseEntity{
			IsActive:  student.IsActive,
			CreatedAt: time.Now().UTC(),
		},
		Status: New,
	}
}

func (m *studentMapper) StudentResponseMapper(student entities.StudentsEntity) dtos.StudentResponseDTO {

	return dtos.StudentResponseDTO{
		StudentID:      student.StudentID,
		FirstName:      student.FirstName,
		Surname:        student.Surname,
		Email:          student.Email,
		Nationality:    student.Nationality,
		DOB:            student.DOB,
		PlaceOfBirth:   student.PlaceOfBirth,
		Sex:            student.Sex,
		Role:           student.Role,
		Status:         student.Status,
		Semester:       student.Semester,
		GraduationDate: student.GraduationDate,

		IsActive:    student.IsActive,
		CreatedAt:   student.CreatedAt,
		IsGraduated: student.IsGraduated,
	}
}
