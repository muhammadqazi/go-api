package services

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/repositories"
)

type CourseServices interface {
	CreateCourse(dtos.CourseCreateDTO) error
	FetchCourse() error
}

type courseServices struct {
	courseMapper     mappers.CourseMapper
	courseRepository repositories.CourseRepository
}

func NewCourseServices(repo repositories.CourseRepository, mapper mappers.CourseMapper) CourseServices {
	return &courseServices{
		courseRepository: repo,
		courseMapper:     mapper,
	}
}

func (s *courseServices) CreateCourse(course dtos.CourseCreateDTO) error {

	courseEntity := s.courseMapper.CourseCreateMapper(course)

	var schedule []dtos.CourseSchedule

	for _, v := range course.CourseSchedule {
		schedule = append(schedule, v)
	}

	return s.courseRepository.InsertCourse(courseEntity, schedule)
}
func (s *courseServices) FetchCourse() error {
	return nil
}
