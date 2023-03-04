package services

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/repositories"
)

type StudentServices interface {
	CreateStudent(dtos.StudentCreateDTO, uint, string) (uint, error)
	FetchLastStudentID() (uint, error)
	FetchStudentByEmail(string) (entities.StudentsEntity, error)
	FetchStudentByID(uint) (entities.StudentsEntity, error)
	CreateTermRegistration(dtos.TermRegistrationDTO, uint) error
	FetchStudentTimetable(uint) ([]dtos.TimetableSchema, error)
	FetchStudentExamSchedule(uint) ([]dtos.ExamScheduleSchema, error)
	FetchStudentAttendance(uint) ([]dtos.StudentAttendanceSchema, error)
	FetchStudentEnrollmentStatus(uint) (bool, error)
	FetchIsEnrollmentExists(uint) (bool, error)
	ModifyStudentPassword(uint, string) error
	CreateForgotPasswordRequest(dtos.ForgotPasswordRequestDTO) error
	VerifyForgotPasswordCode(uint) (entities.StudentPasswordResetsEntity, error)
	ModifyForgotPasswordFlag(uint) error
	RemoveForgotPasswordCode(uint) error
}

type studentServices struct {
	studentMapper     mappers.StudentMapper
	studentRepository repositories.StudentRepository
}

func NewStudentServices(repo repositories.StudentRepository, mapper mappers.StudentMapper) StudentServices {
	return &studentServices{
		studentRepository: repo,
		studentMapper:     mapper,
	}
}

func (s *studentServices) CreateStudent(student dtos.StudentCreateDTO, sid uint, semester string) (uint, error) {
	m := s.studentMapper.StudentCreateMapper(student, sid, semester)
	return s.studentRepository.InsertStudent(m)
}

func (s *studentServices) FetchStudentByEmail(email string) (entities.StudentsEntity, error) {
	return s.studentRepository.QueryStudentByEmail(email)
}

func (s *studentServices) FetchStudentByID(sid uint) (entities.StudentsEntity, error) {
	return s.studentRepository.QueryStudentByID(sid)
}

func (s *studentServices) FetchLastStudentID() (uint, error) {
	return s.studentRepository.QueryLastStudentID()
}

func (s *studentServices) FetchIsEnrollmentExists(sid uint) (bool, error) {
	return s.studentRepository.QueryIsEnrollmentExists(sid)
}

func (s *studentServices) FetchStudentEnrollmentStatus(sid uint) (bool, error) {
	return s.studentRepository.QueryStudentEnrollmentStatus(sid)
}

func (s *studentServices) CreateTermRegistration(registration dtos.TermRegistrationDTO, sid uint) error {

	/*
		get the instructor id from student table and insert into student_course_request table,
		so that the instructor can approve the request. Each student has a supervisor
	*/

	student, err := s.FetchStudentByID(sid)
	if err != nil {
		return err
	}

	supervisorID := student.SupervisorID
	m := s.studentMapper.TermRegistrationMapper(sid, supervisorID)

	courseIDs := registration.CourseIDs
	return s.studentRepository.InsertStudentEnrollment(m, courseIDs)
}

func (s *studentServices) FetchStudentTimetable(sid uint) ([]dtos.TimetableSchema, error) {
	return s.studentRepository.QueryTimetableByStudentID(sid)
}

func (s *studentServices) FetchStudentExamSchedule(sid uint) ([]dtos.ExamScheduleSchema, error) {
	return s.studentRepository.QueryExamScheduleByStudentID(sid)
}

func (s *studentServices) FetchStudentAttendance(sid uint) ([]dtos.StudentAttendanceSchema, error) {
	return s.studentRepository.QueryStudentAttendanceByStudentID(sid)
}

func (s *studentServices) ModifyStudentPassword(sid uint, password string) error {
	return s.studentRepository.UpdateStudentPassword(sid, password)
}

func (s *studentServices) CreateForgotPasswordRequest(request dtos.ForgotPasswordRequestDTO) error {
	m := s.studentMapper.ForgotPasswordRequestMapper(request)
	return s.studentRepository.InsertForgotPasswordRequest(m)
}

func (s *studentServices) VerifyForgotPasswordCode(sid uint) (entities.StudentPasswordResetsEntity, error) {
	return s.studentRepository.QueryForgotPasswordCode(sid)
}

func (s *studentServices) RemoveForgotPasswordCode(sid uint) error {
	return s.studentRepository.DeleteForgotPasswordCode(sid)
}

func (s *studentServices) ModifyForgotPasswordFlag(sid uint) error {
	return s.studentRepository.UpdateForgotPasswordFlag(sid)
}
