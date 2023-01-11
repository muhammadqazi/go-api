package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type AccountingRepository interface {
	InsertAccounts(account entities.AccountsEntity) (uint, error)
}

type accountingConnection struct {
	conn *gorm.DB
}

func NewAccountingRepository(db *gorm.DB) AccountingRepository {
	return &accountingConnection{
		conn: db,
	}
}

func (r *accountingConnection) InsertAccounts(account entities.AccountsEntity) (uint, error) {

	res := r.conn.Create(&account)

	if res.Error != nil {
		return 0, res.Error
	}

	return account.AccountsID, nil
}
