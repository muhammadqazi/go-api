package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type CurriculumRepository interface {
	InsertCurriculum(entities.CurriculumEntity, dtos.CourseCurriculumSchema) error
	QueryCurriculumByDepartmentID(uint) ([]dtos.CurriculumQueryReturnSchema, error)
}

type curriculumConnection struct {
	conn *gorm.DB
}

func NewCurriculumRepository(db *gorm.DB) CurriculumRepository {
	return &curriculumConnection{
		conn: db,
	}
}

func (r *curriculumConnection) InsertCurriculum(curriculum entities.CurriculumEntity, courseCurriculum dtos.CourseCurriculumSchema) error {
	tx := r.conn.Begin()

	if err := tx.Create(&curriculum).Error; err != nil {
		tx.Rollback()
		return err
	}

	var pivot entities.CourseCurriculumEntity
	pivot.CurriculumID = curriculum.CurriculumID
	pivot.CourseLoad = courseCurriculum.CourseLoad
	pivot.CourseID = courseCurriculum.CourseID
	pivot.IsActive = curriculum.IsActive

	if err := tx.Create(&pivot).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (r *curriculumConnection) QueryCurriculumByDepartmentID(departmentID uint) ([]dtos.CurriculumQueryReturnSchema, error) {
	var courses []dtos.CurriculumQueryReturnSchema

	/**================================================================================================
	 * *                                           INFO
	 *
	 *  	 This is the query that is being executed by the below code

			SELECT sc.course_id, sc.course_load, sc.created_at, sc.updated_at,sc.deleted_at,
			cc.curriculum_id, cc.year,cc.semester, cc.department_id,
			d.name AS department_name, d.department_code , d.number_of_years,
			co.course_id, co.code, co.name,co.credits,co.ects,co.practical,co.theoretical
			FROM course_curriculum_entity sc
			JOIN curriculum_entity cc ON sc.curriculum_id = cc.curriculum_id
			JOIN departments_entity d ON cc.department_id = d.department_id
			JOIN courses_entity co ON sc.course_id = co.course_id
			WHERE cc.department_id=2 AND sc.is_active=true;
	 *
	 *
	 *
	 *================================================================================================**/

	err := r.conn.Table("course_curriculum_entity as sc").
		Select("sc.course_id, sc.course_load, sc.created_at, sc.updated_at, sc.deleted_at, cc.curriculum_id, cc.year, cc.semester, cc.department_id, d.name as department_name, d.department_code, d.number_of_years, co.code as code, co.name as name, co.credits as credits, co.ects as ects, co.practical as practical, co.theoretical as theoretical").
		Joins("join curriculum_entity cc on sc.curriculum_id = cc.curriculum_id").
		Joins("join departments_entity d on cc.department_id = d.department_id").
		Joins("join courses_entity co on sc.course_id = co.course_id").
		Where("cc.department_id = ? AND sc.is_active = ?", departmentID, true).
		Scan(&courses).Error

	if err != nil {
		return nil, err
	}

	return courses, nil
}
