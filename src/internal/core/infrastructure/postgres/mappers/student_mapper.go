package mappers

import (
	"time"

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

func (m *studentMapper) StudentCreateMapper(student dtos.StudentCreateDTO, sid uint, semester string) entities.StudentsEntity {

	return entities.StudentsEntity{
		StudentID:    sid,
		FirstName:    student.FirstName,
		Surname:      student.Surname,
		Email:        student.Email,
		Nationality:  student.Nationality,
		DOB:          student.DOB,
		PlaceOfBirth: student.PlaceOfBirth,
		Sex:          student.Sex,
		Password:     student.Password,
		Role:         student.Role,
		Semester:     semester,
		FacultyID:    student.FacultyID,
		BaseEntity: entities.BaseEntity{
			IsActive:  student.IsActive,
			CreatedAt: time.Now().UTC(),
		},
		EnrollmentDate: time.Now().UTC(),
		Status:         New,
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
		EnrollmentDate: student.EnrollmentDate,
		GraduationDate: student.GraduationDate,
		FacultyID:      student.FacultyID,
		PersonalInfoID: student.PersonalInfoID,
		ContactInfoID:  student.ContactInfoID,
		AddressID:      student.AddressID,
		IsActive:       student.IsActive,
		CreatedAt:      student.CreatedAt,
		IsGraduated:    student.IsGraduated,
		IsDeleted:      student.IsDeleted,
	}
}
