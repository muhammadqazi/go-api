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
	BuildingHandler can provide the following services.
	"""
*/

type BuildingHandler interface {
	PostBuilding(c *gin.Context)
	GetBuilding(c *gin.Context)
	GetAllBuildings(c *gin.Context)
}

type buildingHandler struct {
	validator        validation.Validator
	buildingMapper   mappers.BuildingMapper
	buildingServices services.BuildingServices
}

/*
	"""
	This will create a new instance of the BuildingHandler, we will use this as a constructor
	"""
*/

func NewBuildingHandler(service services.BuildingServices, mapper mappers.BuildingMapper, v validation.Validator) BuildingHandler {
	return &buildingHandler{
		buildingMapper:   mapper,
		buildingServices: service,
		validator:        v,
	}
}

func (s *buildingHandler) PostBuilding(c *gin.Context) {
	var building dtos.BuildingCreateDTO

	if err := s.validator.Validate(&building, c); err != nil {
		return
	}

	if err := s.buildingServices.CreateBuilding(building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Building created successfully"})

}
func (s *buildingHandler) GetBuilding(c *gin.Context) {
	code := c.Param("code")

	building, err := s.buildingServices.FetchBuildingByCode(code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": building})
}

func (s *buildingHandler) GetAllBuildings(c *gin.Context) {
	buildings, err := s.buildingServices.FetchAllBuildings()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": buildings})
}
