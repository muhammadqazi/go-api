package repositories

import (
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/SIS-Backend-Go/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type AccountingRepository interface {
	InsertAccounts(entities.AccountsEntity) (uint, error)
	GetAccountByID(uint) (entities.AccountsEntity, error)
	InsertPayment(entities.PaymentsEntity, uint) error
	UpdateAccount(uint, dtos.AccountUpdateDTO) bool
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

func (r *accountingConnection) GetAccountByID(id uint) (entities.AccountsEntity, error) {

	var account entities.AccountsEntity
	err := r.conn.Unscoped().Where("accounts_id = ?", id).First(&account).Error
	return account, err

}

func (r *accountingConnection) InsertPayment(payment entities.PaymentsEntity, sid uint) error {

	var student entities.StudentsEntity

	// Find the student by ID
	if err := r.conn.Where("student_id = ?", sid).First(&student).Error; err != nil {
		return err
	}

	accountID := 1
	payment.AccountID = uint(accountID)

	// Begin a new transaction
	tx := r.conn.Begin()

	// Insert the payment
	if err := tx.Create(&payment).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Get the account
	var account entities.AccountsEntity
	if err := r.conn.Where("accounts_id = ?", accountID).First(&account).Error; err != nil {
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

func (r *accountingConnection) UpdateAccount(id uint, update dtos.AccountUpdateDTO) bool {

	var account entities.AccountsEntity

	if err := r.conn.Model(&account).Where("accounts_id = ?", id).Updates(&update).Error; err != nil {
		return false
	}

	return true
}
