package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/validation"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/services"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"net/http"
)

/*
	"""
	FacultyHandler can provide the following services.
	"""
*/

type FacultyHandler interface {
	PostFaculty(c *gin.Context)
	GetFaculty(c *gin.Context)
}

type facultyHandler struct {
	validator       validation.Validator
	facultyMapper   mappers.FacultyMapper
	facultyServices services.FacultyServices
}

/*
	"""
	This will create a new instance of the FacultyHandler, we will use this as a constructor
	"""
*/

func NewFacultyHandler(service services.FacultyServices, mapper mappers.FacultyMapper, v validation.Validator) FacultyHandler {
	return &facultyHandler{
		facultyMapper:   mapper,
		facultyServices: service,
		validator:       v,
	}
}

func (s *facultyHandler) PostFaculty(c *gin.Context) {
	var faculty dtos.FacultyCreateDTO

	if err := s.validator.Validate(&faculty, c); err != nil {
		return
	}

	if err := s.facultyServices.CreateFaculty(faculty); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Faculty created successfully"})
}
func (s *facultyHandler) GetFaculty(c *gin.Context) {
	var facultyCode = c.Param("code")

	faculty, err := s.facultyServices.FetchFaculty(facultyCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Faculty not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": faculty})
}
