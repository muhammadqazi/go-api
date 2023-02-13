package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type InstructorsRepository interface {
	InsertInstructors(entities.InstructorsEntity) error
	QueryInstructorByEmail(string) (entities.InstructorsEntity, error)
	QueryInstructorByPhone(string) (entities.InstructorsEntity, error)
	QueryTermEnrollmentRequests(uint) ([]dtos.InstructorTermRequests, error)
	UpdateTermEnrollmentRequests(dtos.InstructorApproveEnrollmentRequestDTO) error
}

type instructorsConnection struct {
	conn *gorm.DB
}

func NewInstructorsRepository(db *gorm.DB) InstructorsRepository {
	return &instructorsConnection{
		conn: db,
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

	for _, id := range dto.RequestsIDs {
		if err := r.conn.Table("student_course_request_entity").Where("student_course_request_id = ?", id).Update("is_approved", true).Error; err != nil {
			return err
		}
	}

	return nil
}
