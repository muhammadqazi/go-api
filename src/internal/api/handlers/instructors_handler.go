package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/common/validation"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/services"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"gorm.io/gorm"
	"net/http"
)

/*
	"""
	InstructorsHandler can provide the following services.
	"""
*/

type InstructorsHandler interface {
	CreateInstructors(c *gin.Context)
}

type instructorsHandler struct {
	validator           validation.Validator
	instructorsMapper   mappers.InstructorsMapper
	instructorsServices services.InstructorsServices
}

/*
	"""
	This will create a new instance of the InstructorsHandler, we will use this as a constructor
	"""
*/

func NewInstructorsHandler(service services.InstructorsServices, mapper mappers.InstructorsMapper, v validation.Validator) InstructorsHandler {
	return &instructorsHandler{
		instructorsMapper:   mapper,
		instructorsServices: service,
		validator:           v,
	}
}

func (s *instructorsHandler) CreateInstructors(c *gin.Context) {
	var instructor dtos.InstructorCreateDTO

	if err := s.validator.Validate(&instructor, c); err != nil {
		return
	}

	_, err := s.instructorsServices.FetchInstructorByEmail(instructor.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Email already associated with another instructor"})
		return
	}

	_, err = s.instructorsServices.FetchInstructorByPhone(instructor.PhoneNumber)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Phone number already associated with another instructor"})
		return
	}

	if err := s.instructorsServices.CreateInstructors(instructor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "There was an error performing this action"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": true, "message": "Instructor created successfully"})
}
