package entities

type AccountsEntity struct {
	AccountID     uint    `gorm:"primaryKey;uniqueIndex;not null;autoIncrement" json:"accounts_id"`
	DepartmentFee float32 `gorm:"not null" json:"department_fee"`
	Scholarship   int     `gorm:"not null" json:"scholarship"` // % of scholarship
	Discount      int     `gorm:"default:0" json:"discount"`
	DiscountType  string  `json:"discount_type"`
	Installments  int     `json:"installments"`
	TotalFee      float32 `json:"total_fee"`

	TotalDept       float32 `json:"total_dept"`
	CurrentDept     float32 `json:"current_dept"`
	ApproachingDept float32 `json:"approaching_deadline"`

	StudentID uint `gorm:"column:student_id;not null" json:"student_id"`

	PaymentsEntity []PaymentsEntity `gorm:"foreignkey:AccountID"`
	InvoicesEntity []InvoicesEntity `gorm:"foreignkey:AccountID"`
}
