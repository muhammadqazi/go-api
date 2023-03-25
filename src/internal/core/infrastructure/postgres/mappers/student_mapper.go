package mappers

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/utils"
	"strconv"
	"time"

	"github.com/muhammadqazi/campus-hq-api/src/internal/common/security"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
)

type StudentMapper interface {
	StudentCreateMapper(dtos.StudentCreateDTO, uint, string) entities.StudentsEntity
	StudentResponseMapper(entities.StudentsEntity) dtos.StudentResponseDTO
	TermRegistrationMapper(uint, uint) entities.StudentEnrollmentsEntity
	StudentCourseRequestMapper(uint, uint) entities.StudentCourseRequestEntity
	StudentTimetableMapper([]dtos.TimetableSchema) dtos.TimetableFetchDTO
	StudentExamScheduleMapper([]dtos.ExamScheduleSchema) dtos.ExamScheduleFetchDTO
	StudentCourseAttendanceFetchMapper([]dtos.StudentAttendanceSchema) []dtos.StudentAttendanceFetchDTO
	ForgotPasswordRequestMapper(dtos.ForgotPasswordRequestDTO) entities.StudentPasswordResetsEntity
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

func (m *studentMapper) TermRegistrationMapper(sid uint, supervisorID uint) entities.StudentEnrollmentsEntity {

	semester := utils.GetCurrentSemester()
	year := utils.GetCurrentYear()

	return entities.StudentEnrollmentsEntity{
		StudentID:    sid,
		SupervisorID: supervisorID,
		Semester:     semester,
		Year:         year,
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
	}
}

func (m *studentMapper) StudentTimetableMapper(timetable []dtos.TimetableSchema) dtos.TimetableFetchDTO {
	timetableFetchDTO := dtos.TimetableFetchDTO{}
	timeTableMap := make(map[string][]dtos.LectureInfo)

	for _, t := range timetable {
		lecture := dtos.LectureInfo{
			CourseID:      t.CourseID,
			CourseCode:    t.Code,
			CourseName:    t.Name,
			StartTime:     t.StartTime,
			EndTime:       t.EndTime,
			LectureVenue:  t.LectureVenue,
			Credits:       t.Credits,
			IsTheoretical: t.IsTheoretical,
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

func (m *studentMapper) StudentCourseAttendanceFetchMapper(attendance []dtos.StudentAttendanceSchema) []dtos.StudentAttendanceFetchDTO {
	result := make([]dtos.StudentAttendanceFetchDTO, 0, len(attendance))

	// Group attendance records by course code
	groups := make(map[string][]dtos.StudentAttendanceSchema)
	for _, a := range attendance {
		groups[a.CourseCode] = append(groups[a.CourseCode], a)
	}

	// Calculate attendance statistics for each course
	for code, records := range groups {
		info := dtos.StudentAttendanceFetchDTO{
			CourseCode:  code,
			CourseName:  records[0].CourseName,
			Credits:     records[0].Credits,
			LectureTIme: records[0].LectureTime,
			StartTime:   records[0].StartTime,
			EndTime:     records[0].EndTime,
			Day:         records[0].Day,
		}
		for _, r := range records {
			info.TotalLectures++
			if r.IsAttended {
				info.AttendedLectures++
			}
			if r.IsTheoretical {
				info.TotalTheoreticalLectures++
				if r.IsAttended {
					info.AttendedTheoreticalLectures++
				}
			} else {
				info.TotalPracticalLectures++
				if r.IsAttended {
					info.AttendedPracticalLectures++
				}
			}
		}

		if info.AttendedLectures == 0 || info.AttendedPracticalLectures == 0 || info.AttendedTheoreticalLectures == 0 {
			info.TotalLectureAttendancePercentage = 0
			info.TotalTheoreticalAttendancePercentage = 0
			info.TotalPracticalAttendancePercentage = 0
		} else {
			info.TotalLectureAttendancePercentage = float64(info.AttendedLectures / info.TotalLectures * 100)
			info.TotalTheoreticalAttendancePercentage = float64(info.AttendedTheoreticalLectures / info.TotalTheoreticalLectures * 100)
			info.TotalPracticalAttendancePercentage = float64(info.AttendedPracticalLectures / info.TotalPracticalLectures * 100)
		}
		result = append(result, info)
	}

	return result
}

func (m *studentMapper) ForgotPasswordRequestMapper(request dtos.ForgotPasswordRequestDTO) entities.StudentPasswordResetsEntity {
	code := security.CodeGenerator(6)
	codeInt, _ := strconv.Atoi(code)
	return entities.StudentPasswordResetsEntity{
		StudentID:  request.StudentID,
		ResetCode:  codeInt,
		CreatedAt:  time.Now().UTC(),
		ExpiresAt:  time.Now().UTC().Add(time.Minute * 60),
		IsVerified: false,
	}
}
