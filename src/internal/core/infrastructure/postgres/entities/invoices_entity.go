package entities

import "time"

type InvoicesEntity struct {
	InvoiceID   uint      `gorm:"primaryKey;not null;uniqueIndex" json:"invoice_id"`
	Date        time.Time `gorm:"type:varchar(255);not null" json:"date"`
	Amount      int       `gorm:"not null" json:"amount"`
	Description string    `gorm:"type:varchar(255);not null" json:"description"`
	Installment int       `gorm:"not null" json:"installment"`
	Term        string    `gorm:"type:varchar(255);not null" json:"term"`

	AccountsID uint `gorm:"not null" json:"accounts_id"`
}
