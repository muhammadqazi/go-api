package services

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/mappers"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/repositories"
)

type AccountingServices interface {
	CreateAccounts(dto dtos.AccountCreateDTO) (uint, error)
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

func (s *accountingServices) CreateAccounts(account dtos.AccountCreateDTO) (uint, error) {
	m := s.accountingMapper.AccountsCreateMapper(account)
	return s.accountingRepository.InsertAccounts(m)
}
