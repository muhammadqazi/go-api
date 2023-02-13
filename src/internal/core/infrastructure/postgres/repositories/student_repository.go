package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/utils"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"gorm.io/gorm"
)

type StudentRepository interface {
	InsertStudent(student entities.StudentsEntity) (uint, error)
	QueryLastStudentID() (uint, error)
	QueryStudentByEmail(string) (entities.StudentsEntity, error)
	QueryStudentByID(uint) (entities.StudentsEntity, error)
	InsertStudentEnrollment(entities.StudentEnrollmentsEntity, []uint) error
	QueryTimetableByStudentID(uint) ([]dtos.TimetableSchema, error)
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
		Select("en.student_enrollment_id, en.student_id , en.year, en.semester, req.course_id , req.student_course_request_id, co.name,co.code , sch.day , sch.start_time , sch.end_time , co.credits , sch.lecture_venue").
		Where("en.student_id = ? AND en.semester = ? AND en.year = ? AND req.is_approved", sid, semester, year).
		Scan(&timetable).Error; err != nil {
		return nil, err
	}

	return timetable, nil
}
