package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type InstructorsRepository interface {
	InsertInstructors(entities.InstructorsEntity) error
}

type instructorsConnection struct {
	conn *gorm.DB
}

func NewInstructorsRepository(db *gorm.DB) InstructorsRepository {
	return &instructorsConnection{
		conn: db,
	}
}

func (r *instructorsConnection) InsertInstructors(instructor entities.InstructorsEntity) error {

	return r.conn.Create(&instructor).Error
}
