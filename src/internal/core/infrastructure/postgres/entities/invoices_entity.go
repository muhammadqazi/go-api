package entities

type InvoicesEntity struct {
	InvoiceID   uint   `gorm:"primaryKey;not null;uniqueIndex" json:"invoice_id"`
	Date        string `gorm:"type:varchar(255);not null" json:"date"`
	Amount      int    `gorm:"not null" json:"amount"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	Installment int    `gorm:"not null" json:"installment"`
	Term        string `json:"term"`

	AccountingEntity []AccountingEntity `gorm:"foreignkey:InvoiceID"`
}
