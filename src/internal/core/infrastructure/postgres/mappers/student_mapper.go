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
	TermRegistrationMapper(dtos.TermRegistrationDTO, uint) []entities.StudentEnrollmentsEntity
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
	Registered        = "Registered"
	NotRegistered     = "NotRegistered"
)

const (
	Role = "Student"
)

func (m *studentMapper) StudentCreateMapper(student dtos.StudentCreateDTO, sid uint, semester string) entities.StudentsEntity {

	hashedPassword, _ := security.HashPassword(student.Password)

	return entities.StudentsEntity{
		StudentID:      sid,
		FirstName:      student.FirstName,
		Surname:        student.Surname,
		Email:          student.Email,
		Nationality:    student.Nationality,
		DOB:            student.DOB,
		PlaceOfBirth:   student.PlaceOfBirth,
		Sex:            student.Sex,
		Password:       hashedPassword,
		Role:           Role,
		Semester:       semester,
		DepartmentID:   student.DepartmentID,
		SupervisorID:   student.SupervisorID,
		AcceptanceType: student.AcceptanceType,
		BaseEntity: entities.BaseEntity{
			IsActive:  true,
			CreatedAt: time.Now().UTC(),
		},
		AccessStatus: FullAccess,
		Status:       NotRegistered,
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

func (m *studentMapper) TermRegistrationMapper(registration dtos.TermRegistrationDTO, sid uint) []entities.StudentEnrollmentsEntity {

	var enrollments []entities.StudentEnrollmentsEntity

	for _, courseID := range registration.CourseIDs {
		enrollment := entities.StudentEnrollmentsEntity{
			StudentID:  sid,
			CourseID:   courseID,
			IsApproved: false,
			Semester:   registration.Semester,
			Year:       registration.Year,
			BaseEntity: entities.BaseEntity{
				IsActive:  true,
				CreatedAt: time.Now().UTC(),
			},
		}
		enrollments = append(enrollments, enrollment)
	}

	return enrollments
}
