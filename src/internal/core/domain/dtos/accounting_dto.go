package dtos

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
