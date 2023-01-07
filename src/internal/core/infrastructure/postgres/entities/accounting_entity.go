package entities

type AccountingEntity struct {
	AccountingID uint `gorm:"primaryKey;uniqueIndex;not null" json:"accounting_id"`
	TotalFee     int  `gorm:"not null" json:"total_fee"`
	Scholarship  int  `gorm:"not null" json:"scholarship"`
	Discount     int  `gorm:"not null" json:"discount"`
	Installments int  `gorm:"not null" json:"installment"`

	InvoiceID uint `gorm:"not null" json:"invoice_id"`
	PaymentID uint `gorm:"not null" json:"payment_id"`

	StudentsEntity []StudentsEntity `gorm:"foreignkey:AccountingID"`
}
