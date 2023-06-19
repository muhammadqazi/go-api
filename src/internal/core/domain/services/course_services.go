package services

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/repositories"
)

type CourseServices interface {
	CreateCourse(dtos.CourseCreateDTO) error
	FetchCourseByCourseCode(string) ([]dtos.CourseFetchByCodeSchema, error)
	ModifyCourseByCourseCode(string, dtos.CourseUpdateDTO) error
	RemoveCourseByCourseCode(string) error
	ModifyCourseInstructorByCourseId(dto dtos.CourseInstructorUpdateDTO) error
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

func (s *courseServices) FetchCourseByCourseCode(code string) ([]dtos.CourseFetchByCodeSchema, error) {
	return s.courseRepository.QueryCourseByCourseCode(code)
}

func (s *courseServices) ModifyCourseByCourseCode(code string, course dtos.CourseUpdateDTO) error {
	entity := s.courseMapper.CourseUpdateMapper(course)
	return s.courseRepository.UpdateCourseByCourseCode(code, entity)
}

func (s *courseServices) RemoveCourseByCourseCode(code string) error {
	return s.courseRepository.DeleteCourseByCourseCode(code)
}

func (s *courseServices) ModifyCourseInstructorByCourseId(dto dtos.CourseInstructorUpdateDTO) error {
	return s.courseRepository.UpdateCourseInstructorByCourseId(dto)
}
