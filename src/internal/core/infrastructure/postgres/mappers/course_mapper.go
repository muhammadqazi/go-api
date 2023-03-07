package mappers

import (
	"fmt"
	"time"

	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
)

type CourseMapper interface {
	CourseCreateMapper(dtos.CourseCreateDTO) entities.CoursesEntity
	CourseScheduleMapper(dtos.CourseSchedule, uint) (entities.CourseScheduleEntity, error)
	CourseFetchByCodeMapper([]dtos.CourseFetchByCodeSchema) []dtos.CourseFetchByCodeDTO
	CourseUpdateMapper(dtos.CourseUpdateDTO) entities.CoursesEntity
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
		CourseID:      id,
		Day:           schedule.Day,
		StartTime:     schedule.StartTime,
		EndTime:       schedule.EndTime,
		LectureVenue:  schedule.LectureVenue,
		IsTheoretical: *schedule.IsTheoretical,
	}, nil
}

func (m *courseMapper) CourseFetchByCodeMapper(courses []dtos.CourseFetchByCodeSchema) []dtos.CourseFetchByCodeDTO {
	coursesMap := make(map[string]dtos.CourseFetchByCodeDTO)

	for _, t := range courses {
		if course, ok := coursesMap[t.Code]; ok {
			schedule := dtos.CourseSchedule{
				Day:           t.Day,
				StartTime:     t.StartTime,
				EndTime:       t.EndTime,
				LectureVenue:  t.LectureVenue,
				IsTheoretical: &t.IsTheoretical,
			}

			course.CourseSchedule = append(course.CourseSchedule, schedule)
			coursesMap[t.Code] = course
		} else {
			courseDTO := dtos.CourseFetchByCodeDTO{
				Name:        t.Name,
				Code:        t.Code,
				Description: t.Description,
				Credits:     t.Credits,
				ECTS:        t.Ects,
				Theoretical: t.Theoretical,
				Practical:   t.Practical,
				IsActive:    t.IsActive,
				CourseSchedule: []dtos.CourseSchedule{
					{
						Day:           t.Day,
						StartTime:     t.StartTime,
						EndTime:       t.EndTime,
						LectureVenue:  t.LectureVenue,
						IsTheoretical: &t.IsTheoretical,
					},
				},
			}
			coursesMap[t.Code] = courseDTO
		}
	}

	coursesList := make([]dtos.CourseFetchByCodeDTO, 0, len(coursesMap))
	for _, courseDTO := range coursesMap {
		coursesList = append(coursesList, courseDTO)
	}

	return coursesList
}

func (m *courseMapper) CourseUpdateMapper(dto dtos.CourseUpdateDTO) entities.CoursesEntity {
	return entities.CoursesEntity{
		Name:        dto.Name,
		Code:        dto.Code,
		Description: dto.Description,
		Credits:     dto.Credits,
		ECTS:        dto.ECTS,
		Theoretical: dto.Theoretical,
		Practical:   dto.Practical,
		BaseEntity: entities.BaseEntity{
			UpdatedAt: time.Now().UTC(),
		},
	}
}
