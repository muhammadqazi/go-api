package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/services"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
)

/*
	"""
	StudentHandler can provide the following services.
	"""
*/

type StudentHandler interface {
	CreateStudent(c *gin.Context)
}

type studentHandler struct {
	studentMapper   mappers.StudentMapper
	studentServices services.StudentServices
}

/*
	"""
	This will creates a new instance of the StudentHandler, we will use this as a constructor
	"""
*/

func NewStudentsHandler(service services.StudentServices, mapper mappers.StudentMapper) StudentHandler {
	return &studentHandler{
		studentMapper:   mapper,
		studentServices: service,
	}
}

func (s *studentHandler) CreateStudent(c *gin.Context) {}
