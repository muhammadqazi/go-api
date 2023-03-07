package repositories

import (
	"time"

	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"gorm.io/gorm"
)

type CourseRepository interface {
	InsertCourse(entities.CoursesEntity, []dtos.CourseSchedule) error
	QueryCourseByCourseCode(string) ([]dtos.CourseFetchByCodeSchema, error)
	UpdateCourseByCourseCode(string, entities.CoursesEntity) error
	DeleteCourseByCourseCode(string) error
}

type courseConnection struct {
	conn         *gorm.DB
	courseMapper mappers.CourseMapper
}

func NewCourseRepository(db *gorm.DB, courseMapper mappers.CourseMapper) CourseRepository {
	return &courseConnection{
		conn:         db,
		courseMapper: courseMapper,
	}
}

func (r *courseConnection) InsertCourse(course entities.CoursesEntity, schedule []dtos.CourseSchedule) error {
	tx := r.conn.Begin()

	if err := tx.Create(&course).Error; err != nil {
		tx.Rollback()
		return err
	}

	courseID := course.CourseID
	for _, v := range schedule {
		scheduleEntity, err := r.courseMapper.CourseScheduleMapper(v, courseID)
		if err != nil {
			tx.Rollback()
			return err
		}

		if err = tx.Create(&scheduleEntity).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

func (r *courseConnection) QueryCourseByCourseCode(code string) ([]dtos.CourseFetchByCodeSchema, error) {

	var result []dtos.CourseFetchByCodeSchema
	if err := r.conn.
		Select(`
        co.name,
        co.code,
        co.description,
        co.credits,
        co.ects,
        co.theoretical,
        co.practical,
		co.is_active,
        sch.day,
		sch.is_theoretical,
        sch.start_time,
        sch.end_time,
        sch.lecture_venue`).
		Joins(`
        JOIN course_schedule_entity sch ON sch.course_id = co.course_id`).
		Where("co.code = ?", code).
		Table("courses_entity co").
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *courseConnection) UpdateCourseByCourseCode(code string, course entities.CoursesEntity) error {
	return r.conn.Where("code = ?", code).Updates(&course).Error
}

func (r *courseConnection) DeleteCourseByCourseCode(code string) error {

	update := map[string]interface{}{
		"deleted_at": time.Now().UTC(),
		"is_active":  false,
	}

	return r.conn.Table("courses_entity").Where("code = ?", code).Updates(update).Error
}
