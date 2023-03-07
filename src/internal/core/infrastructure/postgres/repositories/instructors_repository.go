package repositories

import (
	"errors"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/utils"
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
	InsertCourseAttendanceLog(uint) error
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

func (r *instructorsConnection) InsertCourseAttendanceLog(id uint) error {

	tx := r.conn.Begin()

	semester := utils.GetCurrentSemester()
	year := utils.GetCurrentYear()

	var result []dtos.CourseLogSchema
	if err := tx.
		Table("student_enrollments_entity en").
		Joins("JOIN student_course_request_entity req ON req.student_enrollment_id = en.student_enrollment_id").
		Joins("JOIN courses_entity co ON co.course_id = req.course_id").
		Joins("JOIN course_schedule_entity sch ON sch.course_id = req.course_id").
		Select("en.student_id, co.practical AS practical_hours, co.theoretical AS theoretical_hours, co.course_id, co.credits, sch.day, sch.start_time, sch.end_time, sch.is_theoretical").
		Where("en.student_enrollment_id = ? AND en.is_enrolled AND en.semester = ? AND en.year = ?", id, semester, year).
		Scan(&result).
		Error; err != nil {
		tx.Rollback()
		return err
	}

	attendance := entities.StudentAttendanceEntity{
		StudentID: result[0].StudentID,
		Year:      year,
		Semester:  semester,
	}

	var err error
	existingStudent := entities.StudentAttendanceEntity{}
	if err = tx.Where("student_id = ? AND semester = ? AND year = ?", attendance.StudentID, semester, year).First(&existingStudent).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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

	var attendanceLogs []entities.CourseAttendanceEntity

	for _, schedule := range result {
		practicalHours := schedule.PracticalHours
		theoreticalHours := schedule.TheoreticalHours

		totalHours := practicalHours + theoreticalHours
		totalWeeks := totalHours / schedule.Credits

		//theoreticalPerWeek := theoreticalHours / totalWeeks
		//practicalPerWeek := practicalHours / totalWeeks

		semesterStartDate := time.Date(2023, 3, 6, 0, 0, 0, 0, time.UTC)
		semesterStartDay := semesterStartDate.Weekday()

		if schedule.Day != semesterStartDay.String() {
			scheduleWeekday := time.Weekday(0)
			for i := 0; i < 7; i++ {
				if time.Weekday(i).String() == schedule.Day {
					scheduleWeekday = time.Weekday(i)
					break
				}
			}

			diff := int(scheduleWeekday - semesterStartDay)
			if diff < 0 {
				diff += 7
			}

			semesterStartDate = semesterStartDate.AddDate(0, 0, diff)
		}

		for i := 0; i < totalWeeks; i++ {

			/*
				Move this line inside the lectureSlots loop
				semesterStartDate := time.Date(2023, 3, 6, 0, 0, 0, 0, time.UTC)
				semesterStartDay := semesterStartDate.Weekday()
				schedule.StartTime and schedule.EndTime are in string format
				divide them into one hour intervals and insert them into the database
				Convert start and end time strings to time.Time objects
			*/

			startTime, _ := time.Parse("15:04", schedule.StartTime)
			endTime, _ := time.Parse("15:04", schedule.EndTime)

			/* Initialize a slice to store the hourly timestamps */
			var timestamps []time.Time

			/* Initialize a slice to store the lecture slots */
			var lectureSlots [][2]string

			/* Iterate over hours from start to end time */
			for t := startTime; t.Before(endTime); t = t.Add(time.Hour) {
				/* Add the current hour to the timestamp slice */
				timestamps = append(timestamps, t)
				/* Add one hour to the current hour to get the end time of the time slot */
				endOfSlot := t.Add(time.Hour)
				/* If the end of the time slot is after the lecture end time, set it to the lecture end time */
				if endOfSlot.After(endTime) {
					endOfSlot = endTime
				}
				/* Append the start and end times of the time slot to the lectureSlots array */
				lectureSlots = append(lectureSlots, [2]string{t.Format("15:04"), endOfSlot.Format("15:04")})
			}

			for _, slot := range lectureSlots {
				attendanceLogs = append(attendanceLogs, entities.CourseAttendanceEntity{
					StudentAttendanceID: existingStudent.StudentAttendanceID,
					Day:                 schedule.Day,
					StartTime:           slot[0],
					EndTime:             slot[1],
					IsAttended:          false,
					LectureTime:         semesterStartDate,
					CourseID:            schedule.CourseID,
					IsTheoretical:       schedule.IsTheoretical,
				})
			}
			semesterStartDate = semesterStartDate.AddDate(0, 0, 7)
		}
	}

	for _, log := range attendanceLogs {
		if err := tx.Create(&log).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
