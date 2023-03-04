package repositories

import (
	"errors"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"gorm.io/gorm"
	"time"
)

type InstructorsRepository interface {
	InsertInstructors(entities.InstructorsEntity) error
	QueryInstructorByEmail(string) (entities.InstructorsEntity, error)
	QueryInstructorByPhone(string) (entities.InstructorsEntity, error)
	QueryInstructorByID(uint) (entities.InstructorsEntity, error)
	QueryTermEnrollmentRequests(uint) ([]dtos.InstructorTermRequests, error)
	UpdateTermEnrollmentRequests(dtos.InstructorApproveEnrollmentRequestDTO) error
	InsertInstructorCourseEnrollment(entities.InstructorEnrollmentsEntity, dtos.InstructorCourseEnrollmentDTO) error
	QueryInstructorCourseEnrollment(uint) ([]dtos.InstructorEnrollmentsSchema, error)
	UpdateStudentAttendance(entities.StudentAttendanceEntity, dtos.StudentAttendancePatchDTO) error
	QuerySupervisedStudents(uint) ([]dtos.SupervisedStudentSchema, error)
	QueryRegisteredStudentsBySupervisorID(uint) ([]dtos.RegisteredStudentsDTO, error)
}

type instructorsConnection struct {
	conn             *gorm.DB
	instructorMapper mappers.InstructorsMapper
}

func NewInstructorsRepository(db *gorm.DB, instructorMapper mappers.InstructorsMapper) InstructorsRepository {
	return &instructorsConnection{
		conn:             db,
		instructorMapper: instructorMapper,
	}
}

func (r *instructorsConnection) InsertInstructors(instructor entities.InstructorsEntity) error {
	return r.conn.Create(&instructor).Error
}

func (r *instructorsConnection) QueryInstructorByEmail(email string) (entities.InstructorsEntity, error) {
	var instructor entities.InstructorsEntity
	err := r.conn.Unscoped().Where("email = ?", email).First(&instructor).Error
	return instructor, err

}

func (r *instructorsConnection) QueryInstructorByPhone(phone string) (entities.InstructorsEntity, error) {
	var instructor entities.InstructorsEntity
	err := r.conn.Unscoped().Where("phone_number =?", phone).First(&instructor).Error
	return instructor, err
}

func (r *instructorsConnection) QueryInstructorByID(id uint) (entities.InstructorsEntity, error) {
	var instructor entities.InstructorsEntity
	err := r.conn.Unscoped().Where("instructor_id =?", id).First(&instructor).Error
	return instructor, err
}

func (r *instructorsConnection) QueryTermEnrollmentRequests(id uint) ([]dtos.InstructorTermRequests, error) {

	/**================================================================================================
		 * *                                           INFO
		 *
		 *  	This is the query that is being executed by the below code

				SELECT en.student_enrollment_id AS enrollment_id,
				ins.last_name AS supervisor_name, ins.last_name AS supervisor_surname, en.instructor_id AS supervisor_id,
				en.created_at,en.updated_at,en.deleted_at req.is_approved,en.semester,en.year,en.student_id,
				std.first_name AS student_name,std.surname AS student_surname, std.status AS student_status, std.access_status,
	 req.course_id, co.name AS course_name,co.code AS course_code,co.credits AS course_credits,co.is_active AS course_status
				FROM student_course_request_entity req
				JOIN student_enrollments_entity en ON req.student_enrollment_id = en.student_enrollment_id
				JOIN instructors_entity ins ON ins.instructor_id = en.instructor_id
				JOIN students_entity std ON std.student_id = en.student_id
				JOIN courses_entity co ON co.course_id = req.course_id
				WHERE en.is_active=true AND req.is_approved=false AND en.instructor_id = 10;

			 *
		 *
	* * ==============================================================================================*/

	var result []dtos.InstructorTermRequests
	if err := r.conn.Table("student_course_request_entity req").
		Select("en.student_enrollment_id AS enrollment_id,ins.last_name AS supervisor_name, ins.last_name AS supervisor_surname, en.supervisor_id, en.created_at, en.updated_at, en.deleted_at, en.is_enrolled, en.semester, en.year, en.student_id, std.first_name AS student_name, std.surname AS student_surname, std.status AS student_status, std.access_status, req.course_id, co.name AS course_name, co.code AS course_code, co.credits AS course_credits, co.is_active AS course_status, co.ects AS ects, co.practical,co.theoretical").
		Joins("JOIN student_enrollments_entity en ON req.student_enrollment_id = en.student_enrollment_id").
		Joins("JOIN instructors_entity ins ON ins.instructor_id = en.supervisor_id").
		Joins("JOIN students_entity std ON std.student_id = en.student_id").
		Joins("JOIN courses_entity co ON co.course_id = req.course_id").
		Where("en.is_active = ? AND en.is_enrolled = ? AND en.supervisor_id = ?", true, false, id).
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil

}

func (r *instructorsConnection) UpdateTermEnrollmentRequests(dto dtos.InstructorApproveEnrollmentRequestDTO) error {

	if *dto.IsDeclined {
		update := map[string]interface{}{
			"is_enrolled": false,
			"declined_at": time.Now().UTC(),
		}
		return r.conn.Table("student_enrollments_entity").Where("student_enrollment_id = ?", dto.EnrollmentID).Updates(update).Error

	}

	update := map[string]interface{}{
		"is_enrolled": true,
		"approved_at": time.Now().UTC(),
	}

	return r.conn.Table("student_enrollments_entity").Where("student_enrollment_id = ?", dto.EnrollmentID).Updates(update).Error

}

func (r *instructorsConnection) InsertInstructorCourseEnrollment(enrollment entities.InstructorEnrollmentsEntity, courseInfo dtos.InstructorCourseEnrollmentDTO) error {
	tx := r.conn.Begin()

	/*
		First check if the instructor is already enrolled in the enrollments if yes then we will make only
		courses in InstructorCoursesEntity table otherwise we will make a new entry in InstructorEnrollmentsEntity
	*/

	existingEnrollment := entities.InstructorEnrollmentsEntity{}
	var err error
	if err = tx.Where("instructor_id = ?", enrollment.InstructorID).First(&existingEnrollment).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Instructor is not enrolled, create a new enrollment entry
		if err := tx.Create(&enrollment).Error; err != nil {
			tx.Rollback()
			return err
		}
		existingEnrollment = enrollment
	}

	// Create entries for the instructor's courses in the InstructorCoursesEntity table
	for _, courseID := range courseInfo.CourseIDs {
		entity := r.instructorMapper.InstructorCoursesMapper(existingEnrollment.InstructorEnrollmentID, courseID, courseInfo)
		if err := tx.Create(&entity).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *instructorsConnection) QueryInstructorCourseEnrollment(id uint) ([]dtos.InstructorEnrollmentsSchema, error) {

	var result []dtos.InstructorEnrollmentsSchema

	if err := r.conn.
		Model(&entities.InstructorEnrollmentsEntity{}).
		Select(`
			en.created_at AS enrollment_date,
			en.is_active,
			en.instructor_id,
			ins.first_name,
			ins.last_name,
			ins.email,
			ins.is_active AS instructor_status,
			co.course_id,
			sch.day, 
			sch.start_time, 
			sch.end_time, 
			sch.lecture_venue, 
			sch.course_schedule_id,
			co.name,
			co.code,
			co.credits,
			co.theoretical,
			co.practical`).
		Joins(`
			JOIN instructor_courses_entity inco ON en.instructor_enrollment_id = inco.instructor_enrollment_id
			JOIN courses_entity co ON inco.course_id = co.course_id
			JOIN course_schedule_entity sch ON sch.course_id = inco.course_id
			JOIN instructors_entity ins ON ins.instructor_id = en.instructor_id`).
		Where("en.instructor_id = ?", id).
		Table("instructor_enrollments_entity en").
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *instructorsConnection) UpdateStudentAttendance(attendance entities.StudentAttendanceEntity, lecture dtos.StudentAttendancePatchDTO) error {

	tx := r.conn.Begin()

	existingStudent := entities.StudentAttendanceEntity{}
	var err error
	if err = tx.Where("student_id = ? AND semester = ? AND year = ?", attendance.StudentID, attendance.Semester, attendance.Year).First(&existingStudent).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := tx.Create(&attendance).Error; err != nil {
			tx.Rollback()
			return err
		}
		existingStudent = attendance
	}

	attendanceID := existingStudent.StudentAttendanceID
	courseAttendance := r.instructorMapper.CourseAttendancePatchMapper(lecture, attendanceID)

	if err := tx.Create(&courseAttendance).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}

func (r *instructorsConnection) QuerySupervisedStudents(id uint) ([]dtos.SupervisedStudentSchema, error) {
	var doc []dtos.SupervisedStudentSchema
	if err := r.conn.
		Table("students_entity std").
		Select("std.student_id, std.first_name, std.surname, std.email, std.nationality, std.dob, std.sex, std.role, std.status, std.access_status, dep.department_id, dep.department_code, dep.name AS department_name, dep.number_of_years, dep.description AS department_description, fac.faculty_id, fac.name AS faculty_name").
		Joins("JOIN departments_entity dep ON std.department_id = dep.department_id").
		Joins("JOIN faculties_entity fac ON fac.department_id = std.department_id").
		Where("std.supervisor_id = ? AND std.is_active", id).
		Scan(&doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}

func (r *instructorsConnection) QueryRegisteredStudentsBySupervisorID(id uint) ([]dtos.RegisteredStudentsDTO, error) {
	var result []dtos.RegisteredStudentsDTO
	if err := r.conn.
		Select(`
        en.year, 
        en.semester, 
        en.is_enrolled, 
        en.approved_at, 
        en.declined_at,
        CONCAT(std.first_name, ' ', std.surname) AS name, 
        std.student_id, 
        dep.department_code, 
        dep.name AS department_name`).
		Joins(`
        JOIN students_entity std ON en.student_id = std.student_id
        JOIN departments_entity dep ON std.department_id = dep.department_id
        JOIN faculties_entity fac ON fac.department_id = std.department_id`).
		Where("std.supervisor_id = ? AND en.is_enrolled", id).
		Table("student_enrollments_entity en").
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
