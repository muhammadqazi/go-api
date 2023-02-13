package services

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/repositories"
)

type ExamServices interface {
	CreateExamSchedule(dtos.ExamScheduleCreateDTO) error
	FetchExam()
}

type examServices struct {
	examMapper     mappers.ExamMapper
	examRepository repositories.ExamRepository
}

func NewExamServices(repo repositories.ExamRepository, mapper mappers.ExamMapper) ExamServices {
	return &examServices{
		examRepository: repo,
		examMapper:     mapper,
	}
}

func (s *examServices) CreateExamSchedule(exam dtos.ExamScheduleCreateDTO) error {

	entity, err := s.examMapper.ExamScheduleCreateMapper(exam)
	if err != nil {
		return err
	}

	return s.examRepository.InsertExamSchedule(entity)
}

func (s *examServices) FetchExam() {}
