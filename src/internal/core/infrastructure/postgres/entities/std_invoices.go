package entities

import (
	"gorm.io/gorm"
)

type StdInvoicesEntity struct {
	gorm.Model

	StdInvoiceID uint   `gorm:"primary_key;not null;uniqueIndex" json:"std_invoice_id"`
	Date         string `gorm:"type:varchar(255);not null" json:"date"`
	Amount       int    `gorm:"not null" json:"amount"`
	Description  string `gorm:"type:varchar(255);not null" json:"description"`
	Installment  int    `gorm:"not null" json:"installment"`
	Term         string `json:"term"`
}
