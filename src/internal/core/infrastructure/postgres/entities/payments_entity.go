package entities

import "time"

type PaymentsEntity struct {
	PaymentID   uint      `gorm:"primaryKey;not null;uniqueIndex" json:"payment_id"`
	Date        time.Time `gorm:"type:varchar(255);not null" json:"date"`
	Amount      int       `gorm:"not null" json:"amount"`
	ProcessType string    `gorm:"type:varchar(255);not null" json:"process_type"`
	Currency    string    `gorm:"type:varchar(255);not null" json:"currency"`

	BuyRate  float32 `gorm:"not null" json:"buy_rate"`
	SellRate float32 `gorm:"not null" json:"sell_rate"`

	AccountsID uint `gorm:"not null" json:"accounts_id"`
}
