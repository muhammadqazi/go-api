package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type InstructorsRepository interface {
	InsertInstructors(entities.InstructorsEntity) error
	QueryInstructorByEmail(string) (entities.InstructorsEntity, error)
	QueryInstructorByPhone(string) (entities.InstructorsEntity, error)
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

func (r *instructorsConnection) QueryInstructorByEmail(email string) (entities.InstructorsEntity, error) {
	var instructor entities.InstructorsEntity
	err := r.conn.Unscoped().Where("email = ?", email).First(&instructor).Error
	return instructor, err

}

func (r *instructorsConnection) QueryInstructorByPhone(phone string) (entities.InstructorsEntity, error) {
	var instructor entities.InstructorsEntity
	err := r.conn.Unscoped().Where("phone_number =?", phone).First(&instructor).Error
	return instructor, err
}
