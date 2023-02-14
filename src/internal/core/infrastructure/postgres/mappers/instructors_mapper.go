package mappers

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/security"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"strings"
	"time"
)

type InstructorsMapper interface {
	InstructorCreateMapper(dtos.InstructorCreateDTO) entities.InstructorsEntity
	InstructorTermRequestsMapper([]dtos.InstructorTermRequests) []dtos.InstructorTermRequestsFetchDTO
	InstructorCourseEnrollmentMapper(dtos.InstructorCourseEnrollmentDTO) entities.InstructorEnrollmentsEntity
	InstructorCoursesMapper(uint, uint, dtos.InstructorCourseEnrollmentDTO) entities.InstructorCoursesEntity
	InstructorFetchCoursesMapper([]dtos.InstructorEnrollmentsSchema) dtos.InstructorEnrollmentsFetchDTO
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

func (s *instructorsMapper) InstructorCourseEnrollmentMapper(enrollment dtos.InstructorCourseEnrollmentDTO) entities.InstructorEnrollmentsEntity {
	return entities.InstructorEnrollmentsEntity{
		InstructorID: enrollment.InstructorID,
		BaseEntity: entities.BaseEntity{
			CreatedAt: time.Now().UTC(),
			IsActive:  true,
		},
	}
}

func (s *instructorsMapper) InstructorCoursesMapper(enrollmentID uint, courseID uint, info dtos.InstructorCourseEnrollmentDTO) entities.InstructorCoursesEntity {
	return entities.InstructorCoursesEntity{
		InstructorEnrollmentID: enrollmentID,
		CourseID:               courseID,
		Semester:               strings.ToLower(info.Semester),
		Year:                   info.Year,
	}
}

func (s *instructorsMapper) InstructorFetchCoursesMapper(courses []dtos.InstructorEnrollmentsSchema) dtos.InstructorEnrollmentsFetchDTO {
	enrollmentsFetchDTO := dtos.InstructorEnrollmentsFetchDTO{}
	enrollmentsMap := make(map[string][]dtos.CourseEnrollmentInfo)

	for _, t := range courses {
		lecture := dtos.CourseEnrollmentInfo{
			EnrollmentDate:   t.EnrollmentDate,
			EnrollmentStatus: t.IsActive,
			ID:               t.CourseID,
			Name:             t.CourseName,
			Code:             t.CourseCode,
			Credits:          t.Credits,
			Practical:        t.Practical,
			Theoretical:      t.Theoretical,
			Day:              t.Day,
			StartTime:        t.StartTime,
			EndTime:          t.EndTime,
			LectureVenue:     t.LectureVenue,
			CourseScheduleID: t.CourseScheduleID,
		}

		dayLectures, exist := enrollmentsMap[t.Day]
		if !exist {
			dayLectures = []dtos.CourseEnrollmentInfo{}
		}

		dayLectures = append(dayLectures, lecture)
		enrollmentsMap[t.Day] = dayLectures
	}

	enrollmentsFetchDTO.InstructorID = courses[0].InstructorID
	enrollmentsFetchDTO.FirstName = courses[0].FirstName
	enrollmentsFetchDTO.LastName = courses[0].LastName
	enrollmentsFetchDTO.InstructorEmail = courses[0].InstructorEmail
	enrollmentsFetchDTO.InstructorStatus = courses[0].InstructorStatus

	for day, lectures := range enrollmentsMap {
		lecturesInfo := dtos.LecturesEnrollmentInfo{
			Day:     day,
			Courses: lectures,
		}

		enrollmentsFetchDTO.Lectures = append(enrollmentsFetchDTO.Lectures, lecturesInfo)
	}

	return enrollmentsFetchDTO
}
