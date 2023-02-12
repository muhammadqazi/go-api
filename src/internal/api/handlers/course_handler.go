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
	CourseHandler can provide the following services.
	"""
*/

type CourseHandler interface {
	CreateCourse(c *gin.Context)
	GetCourse(c *gin.Context)
}

type courseHandler struct {
	validator      validation.Validator
	courseMapper   mappers.CourseMapper
	courseServices services.CourseServices
}

/*
	"""
	This will create a new instance of the CourseHandler, we will use this as a constructor
	"""
*/

func NewCourseHandler(service services.CourseServices, mapper mappers.CourseMapper, v validation.Validator) CourseHandler {
	return &courseHandler{
		courseMapper:   mapper,
		courseServices: service,
		validator:      v,
	}
}

func (s *courseHandler) CreateCourse(c *gin.Context) {
	var course dtos.CourseCreateDTO

	if err := s.validator.Validate(&course, c); err != nil {
		return
	}

	if err := s.courseServices.CreateCourse(course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Course created successfully"})
}
func (s *courseHandler) GetCourse(c *gin.Context) {}
