package repositories

import (
	"errors"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
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

				SELECT req.student_course_request_id AS request_id,
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
		Select("req.student_course_request_id AS request_id, ins.last_name AS supervisor_name, ins.last_name AS supervisor_surname, en.instructor_id AS supervisor_id, en.created_at, en.updated_at, en.deleted_at, req.is_approved, en.semester, en.year, en.student_id, std.first_name AS student_name, std.surname AS student_surname, std.status AS student_status, std.access_status, req.course_id, co.name AS course_name, co.code AS course_code, co.credits AS course_credits, co.is_active AS course_status, co.ects AS ects, co.practical,co.theoretical").
		Joins("JOIN student_enrollments_entity en ON req.student_enrollment_id = en.student_enrollment_id").
		Joins("JOIN instructors_entity ins ON ins.instructor_id = en.instructor_id").
		Joins("JOIN students_entity std ON std.student_id = en.student_id").
		Joins("JOIN courses_entity co ON co.course_id = req.course_id").
		Where("en.is_active = ? AND req.is_approved = ? AND en.instructor_id = ?", true, false, id).
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil

}

func (r *instructorsConnection) UpdateTermEnrollmentRequests(dto dtos.InstructorApproveEnrollmentRequestDTO) error {

	if *dto.IsDeclined {
		update := map[string]interface{}{
			"is_approved": false,
			"declined_at": time.Now().UTC(),
		}
		return r.conn.Table("student_course_request_entity").Where("student_course_request_id = ?", dto.RequestID).Updates(update).Error

	}

	update := map[string]interface{}{
		"is_approved": true,
		"approved_at": time.Now().UTC(),
	}

	return r.conn.Table("student_course_request_entity").Where("student_course_request_id = ?", dto.RequestID).Updates(update).Error

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
