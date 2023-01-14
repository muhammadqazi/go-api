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
	Amount      int     `json:"amount"`
	ProcessType string  `json:"process_type"`
	Currency    string  `json:"currency"`
	BuyRate     float32 `json:"buy_rate"`
	SellRate    float32 `json:"sell_rate"`
}
type AccountUpdateDTO struct {
	TotalDept int `json:"total_dept"`
}
