package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/validation"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/services"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"net/http"
)

/*
	"""
	ExamHandler can provide the following services.
	"""
*/

type ExamHandler interface {
	PostExamSchedule(c *gin.Context)
	GetExam(c *gin.Context)
}

type examHandler struct {
	validator    validation.Validator
	examMapper   mappers.ExamMapper
	examServices services.ExamServices
}

/*
	"""
	This will create a new instance of the ExamHandler, we will use this as a constructor
	"""
*/

func NewExamHandler(service services.ExamServices, mapper mappers.ExamMapper, v validation.Validator) ExamHandler {
	return &examHandler{
		examMapper:   mapper,
		examServices: service,
		validator:    v,
	}
}

func (s *examHandler) PostExamSchedule(c *gin.Context) {
	var schedule dtos.ExamScheduleCreateDTO

	if err := s.validator.Validate(&schedule, c); err != nil {
		return
	}

	if err := s.examServices.CreateExamSchedule(schedule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Exam Schedule Created Successfully"})
}
func (s *examHandler) GetExam(c *gin.Context) {}
