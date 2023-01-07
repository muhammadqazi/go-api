package entities

import "time"

type PaymentsEntity struct {
	PaymentID   uint      `gorm:"primaryKey;not null;uniqueIndex" json:"payment_id"`
	Date        time.Time `gorm:"type:varchar(255);not null" json:"date"`
	Amount      int       `gorm:"not null" json:"amount"`
	ProcessType string    `gorm:"type:varchar(255);not null" json:"process_type"`

	AccountingEntity []AccountingEntity `gorm:"foreignkey:PaymentID"`
}
