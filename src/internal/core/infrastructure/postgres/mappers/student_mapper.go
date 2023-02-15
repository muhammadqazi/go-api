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
	TermRegistrationMapper(dtos.TermRegistrationDTO, uint, uint) entities.StudentEnrollmentsEntity
	StudentCourseRequestMapper(uint, uint) entities.StudentCourseRequestEntity
	StudentTimetableMapper([]dtos.TimetableSchema) dtos.TimetableFetchDTO
	StudentExamScheduleMapper([]dtos.ExamScheduleSchema) dtos.ExamScheduleFetchDTO
	StudentAttendanceFetchMapper([]dtos.StudentAttendanceSchema) dtos.StudentAttendanceFetchDTO
}

type studentMapper struct {
}

func NewStudentMapper() StudentMapper {
	return &studentMapper{}
}

const (
	New               = "new"
	Active            = "active"
	InActive          = "in-active"
	Pending           = "pending"
	FullAccess        = "full-access"
	ProvisionalAccess = "provisional-access"
	NoAccess          = "no-access"
	Registered        = "registered"
	NotRegistered     = "not-registered"
)

const (
	Role = "student"
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

func (m *studentMapper) TermRegistrationMapper(registration dtos.TermRegistrationDTO, sid uint, supervisorID uint) entities.StudentEnrollmentsEntity {
	return entities.StudentEnrollmentsEntity{
		StudentID:    sid,
		InstructorID: supervisorID,
		Semester:     registration.Semester,
		Year:         registration.Year,
		BaseEntity: entities.BaseEntity{
			IsActive:  true,
			CreatedAt: time.Now().UTC(),
		},
	}
}

func (m *studentMapper) StudentCourseRequestMapper(enrollmentID uint, courseID uint) entities.StudentCourseRequestEntity {
	return entities.StudentCourseRequestEntity{
		CourseID:            courseID,
		StudentEnrollmentID: enrollmentID,
		IsApproved:          false,
	}
}

func (m *studentMapper) StudentTimetableMapper(timetable []dtos.TimetableSchema) dtos.TimetableFetchDTO {
	timetableFetchDTO := dtos.TimetableFetchDTO{}
	timeTableMap := make(map[string][]dtos.LectureInfo)

	for _, t := range timetable {
		lecture := dtos.LectureInfo{
			CourseID:     t.CourseID,
			CourseCode:   t.Code,
			CourseName:   t.Name,
			StartTime:    t.StartTime,
			EndTime:      t.EndTime,
			LectureVenue: t.LectureVenue,
			Credits:      t.Credits,
		}

		dayLectures, ok := timeTableMap[t.Day]
		if !ok {
			dayLectures = []dtos.LectureInfo{}
		}

		dayLectures = append(dayLectures, lecture)
		timeTableMap[t.Day] = dayLectures
	}

	timetableFetchDTO.StudentID = timetable[0].StudentID
	timetableFetchDTO.Year = timetable[0].Year
	timetableFetchDTO.Semester = timetable[0].Semester

	for day, lectures := range timeTableMap {
		timeTableInfo := dtos.TimeTableInfo{
			Day:      day,
			Lectures: lectures,
		}

		timetableFetchDTO.Timetable = append(timetableFetchDTO.Timetable, timeTableInfo)
	}

	return timetableFetchDTO
}

func (m *studentMapper) StudentExamScheduleMapper(examSchedule []dtos.ExamScheduleSchema) dtos.ExamScheduleFetchDTO {
	return dtos.ExamScheduleFetchDTO{
		Schedule: examSchedule,
	}
}

func (m *studentMapper) StudentAttendanceFetchMapper(attendance []dtos.StudentAttendanceSchema) dtos.StudentAttendanceFetchDTO {
	// Initialize the result object
	result := dtos.StudentAttendanceFetchDTO{
		Attendance: make([]dtos.CourseAttendanceInfo, 0, len(attendance)),
	}

	// Map the attendance records to CourseAttendanceInfo objects
	courseAttendanceMap := make(map[uint]*dtos.CourseAttendanceInfo)
	for _, a := range attendance {
		courseID := a.CourseID
		if info, ok := courseAttendanceMap[courseID]; ok {
			// Update the existing CourseAttendanceInfo object
			info.TotalLectures++
			if a.IsAttended {
				info.AttendedLectures++
			} else {
				info.AbsentLectures++
			}
		} else {
			// Create a new CourseAttendanceInfo object
			info := &dtos.CourseAttendanceInfo{
				CourseID:      courseID,
				CourseCode:    a.Code,
				CourseName:    a.Name,
				Credits:       a.Credits,
				TotalLectures: 1,
			}
			if a.IsAttended {
				info.AttendedLectures = 1
			} else {
				info.AbsentLectures = 1
			}
			courseAttendanceMap[courseID] = info
		}
	}

	// Calculate the attendance percentage for each course
	for _, info := range courseAttendanceMap {
		info.PercentageAttendance = int(float64(info.AttendedLectures) / float64(info.TotalLectures) * 100)
		result.Attendance = append(result.Attendance, *info)
	}

	// Set the year and semester fields in the result object
	if len(attendance) > 0 {
		result.Year = attendance[0].Year
		result.Semester = attendance[0].Semester
	}

	return result
}
