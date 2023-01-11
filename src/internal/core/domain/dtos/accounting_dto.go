package dtos

type AccountCreateDTO struct {
	TotalFee     int `json:"total_fee"`
	Scholarship  int `json:"scholarship"`
	Discount     int `json:"discount"`
	Installments int `json:"installments"`
	TotalDept    int `json:"total_dept"`
}
