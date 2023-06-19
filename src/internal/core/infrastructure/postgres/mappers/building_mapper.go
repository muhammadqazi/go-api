package mappers

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
)

type BuildingMapper interface {
	BuildingCreateMapper(dto dtos.BuildingCreateDTO) entities.BuildingsEntity
}

type buildingMapper struct {
}

func NewBuildingMapper() BuildingMapper {
	return &buildingMapper{}
}

func (m *buildingMapper) BuildingCreateMapper(building dtos.BuildingCreateDTO) entities.BuildingsEntity {
	return entities.BuildingsEntity{
		Code:          building.Code,
		Name:          building.Name,
		Description:   building.Description,
		NumberOfRooms: building.NumberOfRooms,
	}
}
