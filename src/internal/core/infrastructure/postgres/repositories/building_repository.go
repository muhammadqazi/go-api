package repositories

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"gorm.io/gorm"
)

type BuildingRepository interface {
	InsertBuilding(entity entities.BuildingsEntity) error
	QueryBuildingByCode(string) (entities.BuildingsEntity, error)
	QueryAllBuildings() ([]entities.BuildingsEntity, error)
}

type buildingConnection struct {
	conn   *gorm.DB
	mapper mappers.BuildingMapper
}

func NewBuildingRepository(db *gorm.DB, mapper mappers.BuildingMapper) BuildingRepository {
	return &buildingConnection{
		conn:   db,
		mapper: mapper,
	}
}

func (r *buildingConnection) InsertBuilding(building entities.BuildingsEntity) error {
	return r.conn.Create(&building).Error
}
func (r *buildingConnection) QueryBuildingByCode(code string) (entities.BuildingsEntity, error) {
	var building entities.BuildingsEntity
	err := r.conn.Where("code = ?", code).First(&building).Error
	return building, err
}

func (r *buildingConnection) QueryAllBuildings() ([]entities.BuildingsEntity, error) {
	var buildings []entities.BuildingsEntity
	err := r.conn.Find(&buildings).Error
	return buildings, err
}
