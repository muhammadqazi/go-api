package entities

import "time"

type StudentPasswordResetsEntity struct {
	ResetID    uint      `gorm:"primaryKey;not null;uniqueIndex" json:"reset_id"`
	StudentID  uint      `gorm:"column:student_id;type:integer;not null"`
	ResetCode  int       `gorm:"column:reset_code;not null"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null"`
	ExpiresAt  time.Time `gorm:"column:expires_at;type:timestamp;not null"`
	IsVerified bool      `gorm:"column:is_verified;type:boolean;not null;default:false"`
}
