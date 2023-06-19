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
	DepartmentHandler can provide the following services.
	"""
*/

type DepartmentHandler interface {
	PostDepartment(c *gin.Context)
	GetDepartment(c *gin.Context)
	GetAllDepartments(c *gin.Context)
}

type departmentHandler struct {
	validator          validation.Validator
	departmentMapper   mappers.DepartmentMapper
	departmentServices services.DepartmentServices
}

/*
	"""
	This will create a new instance of the DepartmentHandler, we will use this as a constructor
	"""
*/

func NewDepartmentHandler(service services.DepartmentServices, mapper mappers.DepartmentMapper, v validation.Validator) DepartmentHandler {
	return &departmentHandler{
		departmentMapper:   mapper,
		departmentServices: service,
		validator:          v,
	}
}

func (s *departmentHandler) PostDepartment(c *gin.Context) {
	var department dtos.DepartmentCreateDTO

	if err := s.validator.Validate(&department, c); err != nil {
		return
	}

	if err := s.departmentServices.CreateDepartment(department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Department created successfully"})

}
func (s *departmentHandler) GetDepartment(c *gin.Context) {
	code := c.Param("code")

	department, err := s.departmentServices.FetchDepartmentByCode(code)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": department})
}

func (s *departmentHandler) GetAllDepartments(c *gin.Context) {
	departments, err := s.departmentServices.FetchAllDepartments()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": departments})
}
