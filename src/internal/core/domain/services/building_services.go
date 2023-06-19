package services

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/repositories"
)

type BuildingServices interface {
	CreateBuilding(dtos.BuildingCreateDTO) error
	FetchBuildingByCode(string) (entities.BuildingsEntity, error)
	FetchAllBuildings() ([]entities.BuildingsEntity, error)
}

type buildingServices struct {
	buildingMapper     mappers.BuildingMapper
	buildingRepository repositories.BuildingRepository
}

func NewBuildingServices(repo repositories.BuildingRepository, mapper mappers.BuildingMapper) BuildingServices {
	return &buildingServices{
		buildingRepository: repo,
		buildingMapper:     mapper,
	}
}

func (s *buildingServices) CreateBuilding(building dtos.BuildingCreateDTO) error {
	entity := s.buildingMapper.BuildingCreateMapper(building)
	return s.buildingRepository.InsertBuilding(entity)
}
func (s *buildingServices) FetchBuildingByCode(code string) (entities.BuildingsEntity, error) {
	return s.buildingRepository.QueryBuildingByCode(code)
}

func (s *buildingServices) FetchAllBuildings() ([]entities.BuildingsEntity, error) {
	return s.buildingRepository.QueryAllBuildings()
}
