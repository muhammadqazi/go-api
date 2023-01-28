package entities

import "time"

type InvoicesEntity struct {
	InvoiceID   uint      `gorm:"primaryKey;not null;uniqueIndex" json:"invoice_id"`
	Date        time.Time `gorm:"type:timestamp;not null" json:"date"`
	Amount      float32   `gorm:"not null" json:"amount"`
	Description string    `gorm:"type:varchar(255);not null" json:"description"`
	Installment int       `gorm:"not null" json:"installment"`
	Term        string    `gorm:"type:varchar(255);not null" json:"term"`

	AccountID uint `gorm:"column:account_id;not null" json:"account_id"`
}
