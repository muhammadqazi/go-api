package mappers

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
)

type AccountsMapper interface {
	AccountsCreateMapper(dtos.AccountCreateDTO) entities.AccountsEntity
}

type accountsMapper struct {
}

func NewAccountingMapper() AccountsMapper {
	return &accountsMapper{}
}

func (m *accountsMapper) AccountsCreateMapper(account dtos.AccountCreateDTO) entities.AccountsEntity {
	return entities.AccountsEntity{
		TotalFee:     account.TotalFee,
		Scholarship:  account.Scholarship,
		Discount:     account.Discount,
		Installments: account.Installments,

		CurrentDept:     account.TotalDept,
		ApproachingDept: 0,
		TotalDept:       account.TotalDept,
	}
}
