package dtos

type AccountCreateDTO struct {
	TotalFee     int    `json:"total_fee"`
	Scholarship  int    `json:"scholarship"`
	Discount     int    `json:"discount"`
	DiscountType string `json:"discount_type"`
	Installments int    `json:"installments"`
	TotalDept    int    `json:"total_dept"`
}

type MakePaymentDTO struct {
	Amount      int    `json:"amount"`
	ProcessType string `json:"process_type"`
	Currency    string `json:"currency"`
	Installment int    `json:"installment"`
}
type AccountUpdateDTO struct {
	TotalDept int `json:"total_dept"`
}
