package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/campus-hq-api/src/internal/common/validation"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/services"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"net/http"
	"strconv"
)

/*
	"""
	CurriculumHandler can provide the following services.
	"""
*/

type CurriculumHandler interface {
	PostCurriculum(*gin.Context)
	GetCurriculumByDepartmentID(*gin.Context)
}

type curriculumHandler struct {
	validator          validation.Validator
	curriculumMapper   mappers.CurriculumMapper
	curriculumServices services.CurriculumServices
}

/*
	"""
	This will create a new instance of the CurriculumHandler, we will use this as a constructor
	"""
*/

func NewCurriculumHandler(service services.CurriculumServices, mapper mappers.CurriculumMapper, v validation.Validator) CurriculumHandler {
	return &curriculumHandler{
		validator:          v,
		curriculumMapper:   mapper,
		curriculumServices: service,
	}
}

func (s *curriculumHandler) PostCurriculum(c *gin.Context) {
	var curriculum dtos.CurriculumCreateDTO

	if err := s.validator.Validate(&curriculum, c); err != nil {
		return
	}

	if err := s.curriculumServices.CreateCurriculum(curriculum); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Curriculum Created Successfully"})
}

func (s *curriculumHandler) GetCurriculumByDepartmentID(c *gin.Context) {

	id := c.Param("id")
	departmentID, _ := strconv.ParseUint(id, 10, 64)

	if doc, err := s.curriculumServices.FetchCurriculumByDepartmentID(uint(departmentID)); err == nil {
		mappedData := s.curriculumMapper.CurriculumFetchMapper(doc)
		c.JSON(http.StatusOK, gin.H{"status": true, "data": mappedData})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Internal Server Error"})
}
