package repositories

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/utils"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"gorm.io/gorm"
)

type StudentRepository interface {
	InsertStudent(student entities.StudentsEntity) (uint, error)
	QueryLastStudentID() (uint, error)
	QueryStudentByEmail(string) (entities.StudentsEntity, error)
	QueryStudentByID(uint) (entities.StudentsEntity, error)
	InsertStudentEnrollment(entities.StudentEnrollmentsEntity, []uint) error
	QueryTimetableByStudentID(uint) ([]dtos.TimetableSchema, error)
	QueryExamScheduleByStudentID(uint) ([]dtos.ExamScheduleSchema, error)
	QueryStudentCourseAttendanceByStudentID(uint) ([]dtos.StudentAttendanceSchema, error)
	QueryStudentEnrollmentStatus(uint) (bool, error)
	QueryIsEnrollmentExists(uint) (bool, error)
	UpdateStudentPassword(uint, string) error
	InsertForgotPasswordRequest(entities.StudentPasswordResetsEntity) error
	QueryForgotPasswordCode(uint) (entities.StudentPasswordResetsEntity, error)
	UpdateForgotPasswordFlag(uint) error
	DeleteForgotPasswordCode(uint) error
}

type studentConnection struct {
	conn          *gorm.DB
	studentMapper mappers.StudentMapper
}

func NewStudentRepository(db *gorm.DB, studentMapper mappers.StudentMapper) StudentRepository {
	return &studentConnection{
		conn:          db,
		studentMapper: studentMapper,
	}
}

func (r *studentConnection) InsertStudent(student entities.StudentsEntity) (uint, error) {

	res := r.conn.Create(&student)

	if res.Error != nil {
		return 0, res.Error
	}

	return student.StudentID, nil
}

func (r *studentConnection) QueryStudentByEmail(email string) (entities.StudentsEntity, error) {

	var student entities.StudentsEntity
	err := r.conn.Unscoped().Where("email = ?", email).First(&student).Error
	return student, err

}

func (r *studentConnection) QueryStudentByID(sid uint) (entities.StudentsEntity, error) {

	var student entities.StudentsEntity
	err := r.conn.Unscoped().Where("student_id = ?", sid).First(&student).Error
	return student, err

}

func (r *studentConnection) QueryLastStudentID() (uint, error) {

	var lastStudent entities.StudentsEntity

	rec := r.conn.Order("student_id DESC").Last(&lastStudent)

	if rec.Error != nil {
		return 0, rec.Error
	}

	return lastStudent.StudentID, nil

}

func (r *studentConnection) QueryIsEnrollmentExists(sid uint) (bool, error) {

	var exists bool

	semester := utils.GetCurrentSemester()
	year := utils.GetCurrentYear()

	if err := r.conn.Table("student_enrollments_entity").
		Where("student_id = ? AND semester = ? AND year = ?", sid, semester, year).
		Select("1").
		Limit(1).
		Scan(&exists).Error; err != nil {
		return true, err
	}

	return exists, nil
}

func (r *studentConnection) QueryStudentEnrollmentStatus(sid uint) (bool, error) {
	var isEnrolled bool

	semester := utils.GetCurrentSemester()
	year := utils.GetCurrentYear()

	if err := r.conn.Table("student_enrollments_entity").
		Where("student_id = ? AND semester = ? AND year = ?", sid, semester, year).
		Pluck("is_enrolled", &isEnrolled).Error; err != nil {
		return false, err
	}

	return isEnrolled, nil
}

func (r *studentConnection) InsertStudentEnrollment(enrollment entities.StudentEnrollmentsEntity, courseIDs []uint) error {

	tx := r.conn.Begin()

	if err := r.conn.Create(&enrollment).Error; err != nil {
		tx.Rollback()
		return err
	}

	enrollmentID := enrollment.StudentEnrollmentID
	for _, cid := range courseIDs {
		studentCourseRequest := r.studentMapper.StudentCourseRequestMapper(enrollmentID, cid)
		if err := r.conn.Create(&studentCourseRequest).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (r *studentConnection) QueryTimetableByStudentID(sid uint) ([]dtos.TimetableSchema, error) {

	var timetable []dtos.TimetableSchema

	/**================================================================================================
	 * *                                           INFO
	 *
	 *  	This is the query that is being executed by the below code

			SELECT en.student_enrollment_id, en.student_id , req.course_id , req.student_course_request_id,
	       	co.name,co.code , sch.day , sch.start_time , sch.end_time , co.credits , sch.lecture_venue, en.year,en.semester
			FROM student_enrollments_entity en
			JOIN student_course_request_entity req ON en.student_enrollment_id = req.student_enrollment_id
			JOIN courses_entity co ON req.course_id = co.course_id
			JOIN course_schedule_entity sch ON req.course_id = sch.course_id
			WHERE en.student_id = 21906778 AND en.semester = 'spring' AND en.year = 2023 AND req.is_approved;
	 *
	 *
	 *
	 *================================================================================================**/

	semester := utils.GetCurrentSemester()
	year := utils.GetCurrentYear()

	if err := r.conn.Table("student_enrollments_entity en").
		Joins("JOIN student_course_request_entity req ON en.student_enrollment_id = req.student_enrollment_id").
		Joins("JOIN courses_entity co ON req.course_id = co.course_id").
		Joins("JOIN course_schedule_entity sch ON req.course_id = sch.course_id").
		Select("en.student_enrollment_id, en.student_id , en.year, en.semester, req.course_id , req.student_course_request_id, co.name,co.code , sch.day , sch.start_time , sch.end_time , co.credits , sch.lecture_venue , sch.is_theoretical").
		Where("en.student_id = ? AND en.semester = ? AND en.year = ? AND en.is_enrolled", sid, semester, year).
		Scan(&timetable).Error; err != nil {
		return nil, err
	}

	return timetable, nil
}

func (r *studentConnection) QueryExamScheduleByStudentID(sid uint) ([]dtos.ExamScheduleSchema, error) {

	var examSchedule []dtos.ExamScheduleSchema

	/**================================================================================================
	 * *                                           INFO
	 *
	 *  	This is the query that is being executed by the below code

			SELECT ex.created_at , ex.is_active , ex.exam_venue , ex.date ,  ex.exam_type, ex.duration , req.course_id , co.code , co.name , co.credits
			FROM student_course_request_entity req
			JOIN student_enrollments_entity en ON en.student_enrollment_id = req.student_enrollment_id
			JOIN exam_schedule_entity ex ON ex.course_id = req.course_id
			JOIN courses_entity co ON co.course_id = req.course_id
			WHERE en.student_id = 21906778 AND en.semester = 'spring' AND en.year = 2023 AND req.is_approved;
	 *
	 *
	 *
	 *================================================================================================**/

	semester := utils.GetCurrentSemester()
	year := utils.GetCurrentYear()

	if err := r.conn.Table("student_course_request_entity req").
		Joins("JOIN student_enrollments_entity en ON en.student_enrollment_id = req.student_enrollment_id").
		Joins("JOIN exam_schedule_entity ex ON ex.course_id = req.course_id").
		Joins("JOIN courses_entity co ON co.course_id = req.course_id").
		Select("ex.created_at, ex.is_active, ex.exam_venue, ex.date, ex.exam_type, ex.duration, req.course_id, co.code, co.name, co.credits").
		Where("en.student_id = ? AND en.semester = ? AND en.year = ? AND req.is_approved = ?", sid, semester, year, true).
		Scan(&examSchedule).Error; err != nil {
		return nil, err
	}
	return examSchedule, nil
}

func (r *studentConnection) QueryStudentCourseAttendanceByStudentID(sid uint) ([]dtos.StudentAttendanceSchema, error) {

	var studentAttendance []dtos.StudentAttendanceSchema

	year := utils.GetCurrentYear()
	semester := utils.GetCurrentSemester()

	if err := r.conn.
		Model(&entities.StudentAttendanceEntity{}).
		Table("student_attendance_entity sa").
		Select(`
        co.lecture_time ,  
		co.day , 
		co.start_time, 
		co.end_time, 
		co.is_attended, 
		co.is_theoretical, 
		cu.code AS course_code , 
		cu.name AS course_name , 
		cu.credits
		`).
		Joins(`
        JOIN course_attendance_entity co ON co.student_attendance_id = sa.student_attendance_id
        JOIN courses_entity cu ON cu.course_id = co.course_id`).
		Where("sa.student_id = ? AND sa.year = ? AND sa.semester = ?", sid, year, semester).
		Scan(&studentAttendance).Error; err != nil {
		return nil, err
	}

	return studentAttendance, nil

}

func (r *studentConnection) UpdateStudentPassword(sid uint, password string) error {

	if err := r.conn.Model(&entities.StudentsEntity{}).
		Where("student_id = ?", sid).
		Update("password", password).Error; err != nil {
		return err
	}

	return nil
}

func (r *studentConnection) InsertForgotPasswordRequest(entity entities.StudentPasswordResetsEntity) error {

	if err := r.conn.Create(&entity).Error; err != nil {
		return err
	}

	return nil
}

func (r *studentConnection) QueryForgotPasswordCode(sid uint) (entities.StudentPasswordResetsEntity, error) {

	var entity entities.StudentPasswordResetsEntity

	if err := r.conn.Model(&entities.StudentPasswordResetsEntity{}).
		Where("student_id = ?", sid).
		First(&entity).Error; err != nil {
		return entity, err
	}

	return entity, nil
}

func (r *studentConnection) DeleteForgotPasswordCode(sid uint) error {

	if err := r.conn.Model(&entities.StudentPasswordResetsEntity{}).
		Where("student_id = ?", sid).
		Delete(&entities.StudentPasswordResetsEntity{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *studentConnection) UpdateForgotPasswordFlag(sid uint) error {

	if err := r.conn.Model(&entities.StudentPasswordResetsEntity{}).
		Where("student_id = ?", sid).
		Update("is_verified", true).Error; err != nil {
		return err
	}

	return nil
}
