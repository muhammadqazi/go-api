package mappers

import (
	"time"

	"github.com/muhammadqazi/campus-hq-api/src/internal/common/security"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/utils"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
)

type InstructorsMapper interface {
	InstructorCreateMapper(dtos.InstructorCreateDTO) entities.InstructorsEntity
	InstructorTermRequestsMapper([]dtos.InstructorTermRequests) []dtos.InstructorTermRequestsFetchDTO
	InstructorCourseEnrollmentMapper(dtos.InstructorCourseEnrollmentDTO) entities.InstructorEnrollmentsEntity
	InstructorCoursesMapper(uint, uint, dtos.InstructorCourseEnrollmentDTO) entities.InstructorCoursesEntity
	InstructorFetchCoursesMapper([]dtos.InstructorEnrollmentsSchema) dtos.InstructorEnrollmentsFetchDTO
	StudentAttendancePatchMapper(dtos.StudentAttendancePatchDTO) entities.StudentAttendanceEntity
	CourseAttendancePatchMapper(dtos.StudentAttendancePatchDTO, uint) entities.CourseAttendanceEntity
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
		Salary:       instructor.Salary,
		OfficeId:     instructor.OfficeId,
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
				EnrollmentID:      request.EnrollmentID,
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
	semester := utils.GetCurrentSemester()
	year := utils.GetCurrentYear()
	return entities.InstructorEnrollmentsEntity{
		InstructorID: enrollment.InstructorID,
		Semester:     semester,
		Year:         year,
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

func (s *instructorsMapper) StudentAttendancePatchMapper(attendance dtos.StudentAttendancePatchDTO) entities.StudentAttendanceEntity {
	year := utils.GetCurrentYear()
	semester := utils.GetCurrentSemester()

	return entities.StudentAttendanceEntity{
		Year:      year,
		Semester:  semester,
		StudentID: attendance.StudentID,
		BaseEntity: entities.BaseEntity{
			CreatedAt: time.Now().UTC(),
			IsActive:  true,
		},
	}
}

func (s *instructorsMapper) CourseAttendancePatchMapper(attendance dtos.StudentAttendancePatchDTO, attendanceID uint) entities.CourseAttendanceEntity {
	return entities.CourseAttendanceEntity{
		StudentAttendanceID: attendanceID,
		CourseID:            attendance.CourseID,
		LectureTime:         attendance.LectureTime,
		IsAttended:          *attendance.IsAttended,
	}
}
