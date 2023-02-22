package dtos

import "time"

type AccountCreateDTO struct {
	TotalFee     int    `json:"total_fee" validate:"required"`
	Scholarship  int    `json:"scholarship" validate:"required"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	Installments int    `json:"installments"`
	TotalDept    int    `json:"total_dept"`
}

type MakePaymentDTO struct {
	Amount      int    `json:"amount" validate:"required"`
	ProcessType string `json:"process_type" validate:"required"`
	Currency    string `json:"currency" validate:"required"`
	Installment int    `json:"installment" validate:"required"`
	StudentID   uint   `json:"student_id" validate:"required"`
}
type AccountUpdateDTO struct {
	TotalDept int `json:"total_dept" validate:"required"`
}

/* Get accounts details */

type AccountDetails struct {
	AccountID          uint      `gorm:"column:account_id"`
	DepartmentFee      float32   `gorm:"column:department_fee"`
	Scholarship        int       `gorm:"column:scholarship"`
	Discount           int       `gorm:"column:discount"`
	DiscountType       string    `gorm:"column:discount_type"`
	Installments       int       `gorm:"column:installments"`
	TotalFee           float32   `gorm:"column:total_fee"`
	TotalDept          float32   `gorm:"column:total_dept"`
	CurrentDept        float32   `gorm:"column:current_dept"`
	ApproachingDept    float32   `gorm:"column:approaching_dept"`
	PaymentDate        time.Time `gorm:"column:payment_date"`
	PaymentCurrency    string    `gorm:"column:payment_currency"`
	InstallmentNumber  int       `gorm:"column:installment_number"`
	InvoiceDate        time.Time `gorm:"column:invoice_date"`
	InvoiceAmount      float32   `gorm:"column:invoice_amount"`
	InvoiceDescription string    `gorm:"column:invoice_description"`
	InvoiceInstallment int       `gorm:"column:invoice_installment"`
	Term               string    `gorm:"column:term"`
}

type PaymentsInfo struct {
	PaymentDate       time.Time `json:"payment_date"`
	PaymentCurrency   string    `json:"payment_currency"`
	InstallmentNumber int       `json:"installment_number"`
}

type InvoicesInfo struct {
	InvoiceDate        time.Time `json:"invoice_date"`
	InvoiceAmount      float32   `json:"invoice_amount"`
	InvoiceDescription string    `json:"invoice_description"`
	InvoiceInstallment int       `json:"invoice_installment"`
	Term               string    `json:"term"`
}

type AccountFetchDTO struct {
	AccountID       uint           `json:"account_id"`
	DepartmentFee   float32        `json:"department_fee"`
	Scholarship     int            `json:"scholarship"`
	Discount        int            `json:"discount"`
	DiscountType    string         `json:"discount_type"`
	Installments    int            `json:"installments"`
	TotalFee        float32        `json:"total_fee"`
	TotalDept       float32        `json:"total_dept"`
	CurrentDept     float32        `json:"current_dept"`
	ApproachingDept float32        `json:"approaching_dept"`
	Payments        []PaymentsInfo `json:"payments"`
	Invoices        []InvoicesInfo `json:"invoices"`
}

/* Accounts Patch DTO */

type AccountPatchDTO struct {
	Scholarship  int    `json:"scholarship"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
}
