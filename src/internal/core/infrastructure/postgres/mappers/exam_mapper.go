package mappers

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/utils"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"time"
)

type ExamMapper interface {
	ExamScheduleCreateMapper(dto dtos.ExamScheduleCreateDTO) (entities.ExamScheduleEntity, error)
	ExamResultsPatchMapper(dto dtos.ExamResultsPatchDTO) entities.ExamResultsEntity
	ExamCourseResultsMapper(dtos.ExamResultsPatchDTO, uint) entities.ExamCourseResultsEntity
}

type examMapper struct {
}

func NewExamMapper() ExamMapper {
	return &examMapper{}
}

func (m *examMapper) ExamScheduleCreateMapper(schedule dtos.ExamScheduleCreateDTO) (entities.ExamScheduleEntity, error) {

	return entities.ExamScheduleEntity{
		Date:      schedule.Date,
		ExamType:  schedule.ExamType,
		Duration:  schedule.Duration,
		ExamVenue: schedule.ExamVenue,
		CourseID:  schedule.CourseID,
		BaseEntity: entities.BaseEntity{
			CreatedAt: time.Now().UTC(),
			IsActive:  true,
		},
	}, nil
}

func (m *examMapper) ExamResultsPatchMapper(results dtos.ExamResultsPatchDTO) entities.ExamResultsEntity {

	semester := utils.GetCurrentSemester()
	year := utils.GetCurrentYear()
	return entities.ExamResultsEntity{
		Semester:  semester,
		Year:      year,
		StudentID: results.StudentID,
		BaseEntity: entities.BaseEntity{
			CreatedAt: time.Now().UTC(),
			IsActive:  true,
		},
	}
}

func (m *examMapper) ExamCourseResultsMapper(results dtos.ExamResultsPatchDTO, examResultsID uint) entities.ExamCourseResultsEntity {
	return entities.ExamCourseResultsEntity{
		ExamResultID: examResultsID,
		CourseID:     results.CourseID,
		Results:      results.Results,
	}
}
