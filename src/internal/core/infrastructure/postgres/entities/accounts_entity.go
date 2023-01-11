package entities

type AccountsEntity struct {
	AccountsID   uint `gorm:"primaryKey;uniqueIndex;not null;autoIncrement" json:"accounts_id"`
	TotalFee     int  `gorm:"not null" json:"total_fee"`
	Scholarship  int  `gorm:"not null" json:"scholarship"`
	Discount     int  `gorm:"not null" json:"discount"`
	Installments int  `json:"installments"`

	CurrentDept     int `json:"current_dept"`
	ApproachingDept int `json:"approaching_dept"`
	TotalDept       int `json:"total_dept"`

	InvoicesEntity []InvoicesEntity `gorm:"foreignkey:AccountsID"`
	PaymentsEntity []PaymentsEntity `gorm:"foreignkey:AccountsID"`

	StudentsEntity []StudentsEntity `gorm:"foreignkey:AccountsID"`
}
