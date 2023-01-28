package entities

import "time"

type PaymentsEntity struct {
	PaymentID   uint      `gorm:"primaryKey;not null;uniqueIndex" json:"payment_id"`
	Date        time.Time `gorm:"type:timestamp;not null" json:"date"`
	Amount      float32   `gorm:"not null" json:"amount"`
	ProcessType string    `gorm:"type:varchar(255);not null" json:"process_type"`
	Currency    string    `gorm:"type:varchar(255);not null" json:"currency"`
	Installment int       `gorm:"not null" json:"installment"`

	AccountID uint `gorm:"column:account_id;not null" json:"account_id"`
}
