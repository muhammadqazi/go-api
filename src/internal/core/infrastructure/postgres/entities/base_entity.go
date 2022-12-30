package entities

type BaseEntity struct {
	CreatedAt string `gorm:"type:varchar(255);not null" json:"created_at"`
	UpdatedAt string `gorm:"type:varchar(255)" json:"updated_at"`
	DeletedAt string `gorm:"type:varchar(255)" json:"deleted_at"`
	IsActive  bool   `gorm:"type:bool;not null" json:"is_active"`
}
