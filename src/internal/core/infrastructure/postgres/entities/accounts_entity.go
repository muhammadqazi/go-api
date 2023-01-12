package entities

type AccountsEntity struct {
	AccountsID   uint   `gorm:"primaryKey;uniqueIndex;not null;autoIncrement" json:"accounts_id"`
	TotalFee     int    `gorm:"not null" json:"total_fee"`
	Scholarship  int    `gorm:"not null" json:"scholarship"` // % of scholarship
	Discount     int    `gorm:"default:0" json:"discount" `
	DiscountType string `json:"discount_type"`
	Installments int    `json:"installments"`

	TotalDept int `json:"total_dept"`

	InvoicesEntity []InvoicesEntity `gorm:"foreignkey:AccountsID"`
	PaymentsEntity []PaymentsEntity `gorm:"foreignkey:AccountsID"`

	StudentsEntity []StudentsEntity `gorm:"foreignkey:AccountsID"`
}
