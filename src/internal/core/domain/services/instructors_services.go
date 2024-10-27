package services

import (
	"fmt"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/repositories"
)

type InstructorsServices interface {
	CreateInstructors(dtos.InstructorCreateDTO) error
	FetchInstructorByEmail(string) (entities.InstructorsEntity, error)
	FetchInstructorByPhone(string) (entities.InstructorsEntity, error)
	FetchTermEnrollmentRequests(uint) ([]dtos.InstructorTermRequests, error)
	ModifyTermEnrollmentRequests(dtos.InstructorApproveEnrollmentRequestDTO) error
	CreateInstructorCourseEnrollment(dtos.InstructorCourseEnrollmentDTO) error
	FetchInstructorCourseEnrollment(uint) ([]dtos.InstructorEnrollmentsSchema, error)
	ModifyStudentAttendance(dto dtos.StudentAttendancePatchDTO) error
	FetchSupervisedStudents(uint, string) ([]dtos.SupervisedStudentSchema, error)
	FetchRegisteredStudentsBySupervisorID(uint) ([]dtos.RegisteredStudentsDTO, error)
	CreateCourseAttendanceLog(uint) error
	FetchAllStudents() ([]dtos.StudentsFetchDTO, error)
	ModifyPassword(dtos.ResetPasswordDTO, uint) error
}

type instructorsServices struct {
	instructorsMapper     mappers.InstructorsMapper
	instructorsRepository repositories.InstructorsRepository
}

func NewInstructorsServices(repo repositories.InstructorsRepository, mapper mappers.InstructorsMapper) InstructorsServices {
	return &instructorsServices{
		instructorsRepository: repo,
		instructorsMapper:     mapper,
	}
}

func (s *instructorsServices) CreateInstructors(instructor dtos.InstructorCreateDTO) error {
	m := s.instructorsMapper.InstructorCreateMapper(instructor)
	return s.instructorsRepository.InsertInstructors(m)
}

func (s *instructorsServices) FetchAllStudents() ([]dtos.StudentsFetchDTO, error) {
	docs, err := s.instructorsRepository.QueryAllStudents()

	if err != nil {
		return nil, fmt.Errorf("error while fetching all students: %w", err)
	}

	fmt.Println(docs)
	return nil, nil
}

func (s *instructorsServices) FetchInstructorByEmail(email string) (entities.InstructorsEntity, error) {
	return s.instructorsRepository.QueryInstructorByEmail(email)
}

func (s *instructorsServices) FetchInstructorByPhone(phone string) (entities.InstructorsEntity, error) {
	return s.instructorsRepository.QueryInstructorByPhone(phone)
}

func (s *instructorsServices) FetchTermEnrollmentRequests(id uint) ([]dtos.InstructorTermRequests, error) {
	return s.instructorsRepository.QueryTermEnrollmentRequests(id)
}

func (s *instructorsServices) ModifyTermEnrollmentRequests(request dtos.InstructorApproveEnrollmentRequestDTO) error {
	return s.instructorsRepository.UpdateTermEnrollmentRequests(request)
}

func (s *instructorsServices) CreateInstructorCourseEnrollment(enrollment dtos.InstructorCourseEnrollmentDTO) error {
	entity := s.instructorsMapper.InstructorCourseEnrollmentMapper(enrollment)
	return s.instructorsRepository.InsertInstructorCourseEnrollment(entity, enrollment)
}

func (s *instructorsServices) FetchInstructorCourseEnrollment(id uint) ([]dtos.InstructorEnrollmentsSchema, error) {
	return s.instructorsRepository.QueryInstructorCourseEnrollment(id)
}

func (s *instructorsServices) ModifyStudentAttendance(attendance dtos.StudentAttendancePatchDTO) error {
	entity := s.instructorsMapper.StudentAttendancePatchMapper(attendance)
	return s.instructorsRepository.UpdateStudentAttendance(entity, attendance)
}

func (s *instructorsServices) FetchSupervisedStudents(id uint, role string) ([]dtos.SupervisedStudentSchema, error) {
	return s.instructorsRepository.QuerySupervisedStudents(id, role)
}

func (s *instructorsServices) FetchRegisteredStudentsBySupervisorID(id uint) ([]dtos.RegisteredStudentsDTO, error) {
	return s.instructorsRepository.QueryRegisteredStudentsBySupervisorID(id)
}

func (s *instructorsServices) CreateCourseAttendanceLog(id uint) error {
	return s.instructorsRepository.InsertCourseAttendanceLog(id)
}

func (s *instructorsServices) ModifyPassword(dto dtos.ResetPasswordDTO, id uint) error {
	return s.instructorsRepository.UpdatePassword(dto, id)
}
