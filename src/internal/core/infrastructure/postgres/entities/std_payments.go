package entities

import (
	"gorm.io/gorm"
)

type StdPaymentsEntity struct {
	gorm.Model

	StdPaymentID uint   `gorm:"primary_key;not null;uniqueIndex" json:"std_payment_id"`
	Date         string `gorm:"type:varchar(255);not null" json:"date"`
	Amount       int    `gorm:"not null" json:"amount"`
	ProcessType  string `gorm:"type:varchar(255);not null" json:"process_type"`
}
