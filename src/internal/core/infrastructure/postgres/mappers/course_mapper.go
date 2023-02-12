package mappers

import (
	"fmt"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"time"
)

type CourseMapper interface {
	CourseCreateMapper(dtos.CourseCreateDTO) entities.CoursesEntity
	CourseScheduleMapper(dtos.CourseSchedule, uint) (entities.CourseScheduleEntity, error)
}

type courseMapper struct {
}

func NewCourseMapper() CourseMapper {
	return &courseMapper{}
}

func (m *courseMapper) CourseCreateMapper(dto dtos.CourseCreateDTO) entities.CoursesEntity {
	return entities.CoursesEntity{
		Name:        dto.Name,
		Code:        dto.Code,
		Description: dto.Description,
		Credits:     dto.Credits,
		ECTS:        dto.ECTS,
		Theoretical: dto.Theoretical,
		Practical:   dto.Practical,
	}
}

func (m *courseMapper) CourseScheduleMapper(schedule dtos.CourseSchedule, id uint) (entities.CourseScheduleEntity, error) {

	/* Just checking if the time is correct*/

	startTime, err := time.Parse("15:04", schedule.StartTime)
	if err != nil {
		return entities.CourseScheduleEntity{}, err
	}

	endTime, err := time.Parse("15:04", schedule.EndTime)
	if err != nil {
		return entities.CourseScheduleEntity{}, err
	}

	if endTime.Before(startTime) || startTime.Equal(endTime) {
		return entities.CourseScheduleEntity{}, fmt.Errorf("end time should be greater than start time")
	}

	return entities.CourseScheduleEntity{
		CourseID:     id,
		Day:          schedule.Day,
		StartTime:    schedule.StartTime,
		EndTime:      schedule.EndTime,
		LectureVenue: schedule.LectureVenue,
	}, nil
}
