package mappers

import (
	"time"

	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
)

type AccountsMapper interface {
	AccountsCreateMapper(dtos.StudentCreateDTO) entities.AccountsEntity
	MakePaymentMapper(dtos.MakePaymentDTO) entities.PaymentsEntity
}

type accountsMapper struct {
}

func NewAccountingMapper() AccountsMapper {
	return &accountsMapper{}
}

func (m *accountsMapper) AccountsCreateMapper(student dtos.StudentCreateDTO) entities.AccountsEntity {

	DiscountType := "none"
	if student.Discount > 0 {
		DiscountType = student.DiscountType
	}

	totalFee := 3500
	scholarship := student.Scholarship
	currentDept := totalFee - (totalFee * scholarship / 100)

	if student.Discount > 0 {
		currentDept -= (currentDept * student.Discount / 100)
	}

	return entities.AccountsEntity{
		TotalFee:     totalFee,
		Scholarship:  scholarship,
		Discount:     student.Discount,
		DiscountType: DiscountType,
		Installments: 2, // default

		TotalDept: currentDept,
	}
}

func (m *accountsMapper) MakePaymentMapper(payment dtos.MakePaymentDTO) entities.PaymentsEntity {

	return entities.PaymentsEntity{
		Amount:      payment.Amount,
		ProcessType: payment.ProcessType,
		Date:        time.Now().UTC(),
		Currency:    payment.Currency,
		BuyRate:     payment.BuyRate,
		SellRate:    payment.SellRate,
	}
}
