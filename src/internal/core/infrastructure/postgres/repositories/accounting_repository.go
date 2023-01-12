package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type AccountingRepository interface {
	InsertAccounts(entities.AccountsEntity) (uint, error)
	InsertPayment(entities.PaymentsEntity, uint) bool
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

func (r *accountingConnection) InsertPayment(payment entities.PaymentsEntity, sid uint) bool {

	var student entities.StudentsEntity

	if err := r.conn.Where("student_id = ?", sid).First(&student).Error; err != nil {
		return false
	}

	payment.AccountsID = student.AccountsID

	if err := r.conn.Create(&payment).Error; err != nil {
		return false
	}

	return true
}
