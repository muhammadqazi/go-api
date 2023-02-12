package entities

import "time"

type BaseEntity struct {
	CreatedAt time.Time `gorm:"type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:null" json:"updated_at"`
	DeletedAt time.Time `gorm:"type:timestamp;default:null" json:"deleted_at"`
	IsActive  bool      `gorm:"type:bool;not null;default:true" json:"is_active"`
}
