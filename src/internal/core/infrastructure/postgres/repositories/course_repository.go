package repositories

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"gorm.io/gorm"
)

type CourseRepository interface {
	InsertCourse(entities.CoursesEntity, []dtos.CourseSchedule) error
	QueryCourse() error
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
func (r *courseConnection) QueryCourse() error {
	return nil
}
