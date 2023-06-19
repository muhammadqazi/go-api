package repositories

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/mappers"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	InsertDepartment(entity entities.DepartmentsEntity) error
	QueryDepartmentById(code string) (entities.DepartmentsEntity, error)
	QueryAllDepartments() ([]entities.DepartmentsEntity, error)
}

type departmentConnection struct {
	conn   *gorm.DB
	mapper mappers.DepartmentMapper
}

func NewDepartmentRepository(db *gorm.DB, mapper mappers.DepartmentMapper) DepartmentRepository {
	return &departmentConnection{
		conn:   db,
		mapper: mapper,
	}
}

func (r *departmentConnection) InsertDepartment(department entities.DepartmentsEntity) error {
	return r.conn.Create(&department).Error
}
func (r *departmentConnection) QueryDepartmentById(code string) (entities.DepartmentsEntity, error) {
	var department entities.DepartmentsEntity
	err := r.conn.Where("department_code = ?", code).First(&department).Error
	return department, err
}

func (r *departmentConnection) QueryAllDepartments() ([]entities.DepartmentsEntity, error) {
	var departments []entities.DepartmentsEntity
	err := r.conn.Find(&departments).Error
	return departments, err
}
