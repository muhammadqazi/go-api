package services

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/repositories"
)

type AccountingServices interface {
	CreateAccounts(dtos.StudentCreateDTO, uint) (uint, error)
	CreatePayment(dtos.MakePaymentDTO, uint) error
}

type accountingServices struct {
	accountingMapper     mappers.AccountsMapper
	accountingRepository repositories.AccountingRepository
}

func NewAccountingServices(repo repositories.AccountingRepository, mapper mappers.AccountsMapper) AccountingServices {
	return &accountingServices{
		accountingRepository: repo,
		accountingMapper:     mapper,
	}
}

func (s *accountingServices) CreateAccounts(account dtos.StudentCreateDTO, sid uint) (uint, error) {
	m := s.accountingMapper.AccountsCreateMapper(account, sid)
	return s.accountingRepository.InsertAccounts(m)
}

func (s *accountingServices) CreatePayment(payment dtos.MakePaymentDTO, sid uint) error {

	m := s.accountingMapper.MakePaymentMapper(payment)

	return s.accountingRepository.InsertPayment(m, sid)
}
