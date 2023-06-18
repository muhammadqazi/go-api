package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/validation"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/services"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
)

/*
	"""
	DepartmentHandler can provide the following services.
	"""
*/

type DepartmentHandler interface {
	PostDepartment(c *gin.Context)
	GetDepartment(c *gin.Context)
}

type departmentHandler struct {
  validator            validation.Validator
	departmentMapper   mappers.DepartmentMapper
	departmentServices services.DepartmentServices
}

/*
	"""
	This will create a new instance of the DepartmentHandler, we will use this as a constructor
	"""
*/

func NewDepartmentHandler(service services.DepartmentServices, mapper mappers.DepartmentMapper,v validation.Validator) DepartmentHandler {
	return &departmentHandler{
		departmentMapper:   mapper,
		departmentServices: service,
		validator:            v,
	}
}

func (s *departmentHandler) PostDepartment(c *gin.Context) {}
func (s *departmentHandler) GetDepartment(c *gin.Context) {}
