package mappers

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/security"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"time"
)

type InstructorsMapper interface {
	InstructorCreateMapper(dtos.InstructorCreateDTO) entities.InstructorsEntity
	InstructorTermRequestsMapper([]dtos.InstructorTermRequests) []dtos.InstructorTermRequestsFetchDTO
}

type instructorsMapper struct {
}

func NewInstructorsMapper() InstructorsMapper {
	return &instructorsMapper{}
}

func (s *instructorsMapper) InstructorCreateMapper(instructor dtos.InstructorCreateDTO) entities.InstructorsEntity {

	hashedPassword, _ := security.HashPassword(instructor.Password)

	return entities.InstructorsEntity{
		FirstName:    instructor.FirstName,
		LastName:     instructor.LastName,
		Email:        instructor.Email,
		Password:     hashedPassword,
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

func (s *instructorsMapper) InstructorTermRequestsMapper(requests []dtos.InstructorTermRequests) []dtos.InstructorTermRequestsFetchDTO {
	studentInfoMap := make(map[uint]dtos.InstructorTermRequestsFetchDTO)

	for _, request := range requests {
		studentID := request.StudentID
		studentInfo, exists := studentInfoMap[studentID]
		if !exists {
			studentInfo = dtos.InstructorTermRequestsFetchDTO{
				SupervisorID:      request.SupervisorID,
				SupervisorName:    request.SupervisorName,
				SupervisorSurname: request.SupervisorSurname,
				StudentID:         request.StudentID,
				StudentName:       request.StudentName,
				StudentSurname:    request.StudentSurname,
				StudentStatus:     request.StudentStatus,
				AccessStatus:      request.AccessStatus,
				Semester:          request.Semester,
				Year:              request.Year,
				IsApproved:        request.IsApproved,
				Courses:           []dtos.CourseApprovalInfo{},
			}
		}
		studentInfo.Courses = append(studentInfo.Courses, dtos.CourseApprovalInfo{
			ID:          request.CourseID,
			Name:        request.CourseName,
			Code:        request.CourseCode,
			Credits:     request.CourseCredits,
			Ects:        request.ECTS,
			Theoretical: request.Theoretical,
			Practical:   request.Practical,
			IsApproved:  request.IsApproved,
			RequestID:   request.RequestID,
		})
		studentInfoMap[studentID] = studentInfo
	}

	var studentInfoArray []dtos.InstructorTermRequestsFetchDTO
	for _, value := range studentInfoMap {
		studentInfoArray = append(studentInfoArray, value)
	}

	return studentInfoArray
}
