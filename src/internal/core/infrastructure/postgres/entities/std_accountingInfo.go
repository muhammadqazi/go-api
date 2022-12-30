package entities

import "gorm.io/gorm"

type StdAccountingInfoEntity struct {
	gorm.Model

	StdAccountingInfoID uint              `gorm:"primary_key;not null;uniqueIndex" json:"std_accounting_info_id"`
	TotalFee            float64           `gorm:"type:varchar(255);not null;uniqueIndex" json:"total_fee"`
	Scholarship         int               `gorm:"not null;uniqueIndex" json:"scholarship"`
	Discount            int               `gorm:"not null;uniqueIndex" json:"discount"`
	Installments        int               `gorm:"not null;uniqueIndex" json:"installment"`
	StdPaymentsEntity   StdPaymentsEntity `gorm:"foreignKey:StdPaymentsID;references:StdPaymentID"`
	StdPaymentsID       uint
	StdInvoicesEntity   StdInvoicesEntity `gorm:"foreignKey:StdInvoicesID;references:StdInvoiceID"`
	StdInvoicesID       uint
}
