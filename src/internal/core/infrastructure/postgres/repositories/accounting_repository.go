package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type AccountingRepository interface {
	InsertAccounts(entities.AccountsEntity) (uint, error)
	GetAccountByStudentID(uint) (entities.AccountsEntity, error)
	InsertPayment(entities.PaymentsEntity, uint) error
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

	return account.AccountID, nil
}

func (r *accountingConnection) GetAccountByStudentID(id uint) (entities.AccountsEntity, error) {

	var account entities.AccountsEntity
	err := r.conn.Unscoped().Where("student_id = ?", id).First(&account).Error
	return account, err

}

func (r *accountingConnection) InsertPayment(payment entities.PaymentsEntity, sid uint) error {

	account, err := r.GetAccountByStudentID(sid)
	if err != nil {
		return err
	}

	accountID := account.AccountID
	payment.AccountID = accountID

	// Begin a new transaction
	tx := r.conn.Begin()

	// Insert the payment
	if err := tx.Create(&payment).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update the account total debt
	account.TotalDept -= payment.Amount
	if err := tx.Save(&account).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	tx.Commit()

	return nil
}
