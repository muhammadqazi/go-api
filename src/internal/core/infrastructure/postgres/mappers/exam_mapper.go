package mappers

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"time"
)

type ExamMapper interface {
	ExamScheduleCreateMapper(dto dtos.ExamScheduleCreateDTO) (entities.ExamScheduleEntity, error)
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
