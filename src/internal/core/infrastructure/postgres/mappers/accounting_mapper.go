package mappers

import (
	"time"

	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
)

type AccountsMapper interface {
	AccountsCreateMapper(dtos.StudentCreateDTO, uint) entities.AccountsEntity
	MakePaymentMapper(dtos.MakePaymentDTO) entities.PaymentsEntity
	AccountsFetchMapper([]dtos.AccountDetails) dtos.AccountFetchDTO
}

type accountsMapper struct {
}

func NewAccountingMapper() AccountsMapper {
	return &accountsMapper{}
}

func (m *accountsMapper) AccountsCreateMapper(student dtos.StudentCreateDTO, sid uint) entities.AccountsEntity {

	DiscountType := "none"
	if student.Discount > 0 {
		DiscountType = student.DiscountType
	}

	totalFee := 3500

	scholarship := student.Scholarship
	currentDept := totalFee - (totalFee * scholarship / 100)

	if student.Discount > 0 {
		currentDept -= currentDept * student.Discount / 100
	}

	installments := 0
	totalDept := currentDept

	if student.Installments > 0 {
		installments = student.Installments
		totalDept = currentDept / installments
	}

	return entities.AccountsEntity{
		DepartmentFee: float32(totalFee),
		Scholarship:   scholarship,
		Discount:      student.Discount,
		DiscountType:  DiscountType,
		Installments:  installments,
		TotalFee:      float32(currentDept),
		StudentID:     sid,

		TotalDept: float32(totalDept),
	}
}

func (m *accountsMapper) MakePaymentMapper(payment dtos.MakePaymentDTO) entities.PaymentsEntity {

	return entities.PaymentsEntity{
		Amount:      float32(payment.Amount),
		ProcessType: payment.ProcessType,
		Date:        time.Now().UTC(),
		Currency:    payment.Currency,
		Installment: payment.Installment,
	}
}

func (m *accountsMapper) AccountsFetchMapper(details []dtos.AccountDetails) dtos.AccountFetchDTO {

	var fetchDto dtos.AccountFetchDTO

	payments := make([]dtos.PaymentsInfo, 0)
	invoices := make([]dtos.InvoicesInfo, 0)

	for _, payment := range details {
		payments = append(payments, dtos.PaymentsInfo{
			PaymentDate:       payment.PaymentDate,
			PaymentCurrency:   payment.PaymentCurrency,
			InstallmentNumber: payment.InstallmentNumber,
		})
	}

	for _, invoice := range details {
		invoices = append(invoices, dtos.InvoicesInfo{
			InvoiceDate:        invoice.InvoiceDate,
			InvoiceAmount:      invoice.InvoiceAmount,
			InvoiceDescription: invoice.InvoiceDescription,
			InvoiceInstallment: invoice.InvoiceInstallment,
			Term:               invoice.Term,
		})
	}

	fetchDto.AccountID = details[0].AccountID
	fetchDto.DepartmentFee = details[0].DepartmentFee
	fetchDto.Scholarship = details[0].Scholarship
	fetchDto.Discount = details[0].Discount
	fetchDto.DiscountType = details[0].DiscountType
	fetchDto.Installments = details[0].Installments
	fetchDto.TotalFee = details[0].TotalFee
	fetchDto.TotalDept = details[0].TotalDept
	fetchDto.CurrentDept = details[0].CurrentDept
	fetchDto.ApproachingDept = details[0].ApproachingDept
	fetchDto.Payments = payments
	fetchDto.Invoices = invoices

	return fetchDto
}
