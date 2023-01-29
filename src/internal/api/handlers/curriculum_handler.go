package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/services"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
)

/*
	"""
	CurriculumHandler can provide the following services.
	"""
*/

type CurriculumHandler interface {
	CreateCurriculum(c *gin.Context)
}

type curriculumHandler struct {
	curriculumMapper   mappers.CurriculumMapper
	curriculumServices services.CurriculumServices
}

/*
	"""
	This will creates a new instance of the CurriculumHandler, we will use this as a constructor
	"""
*/

func NewCurriculumHandler(service services.CurriculumServices, mapper mappers.CurriculumMapper) CurriculumHandler {
	return &curriculumHandler{
		curriculumMapper:   mapper,
		curriculumServices: service,
	}
}

func (s *curriculumHandler) CreateCurriculum(c *gin.Context) {}
