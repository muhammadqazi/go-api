package repositories

import (
  "github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
  "gorm.io/gorm"
)

type DepartmentRepository interface {
    InsertDepartment(entities.DepartmentEntity) error
    QueryDepartment() error
}

type departmentConnection struct {
    conn *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
    return &departmentConnection{
        conn: db,
    }
}

func (r *departmentConnection) InsertDepartment(department entities.DepartmentEntity) error {}
func (r *departmentConnection) QueryDepartment() error {}

