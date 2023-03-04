package repositories

import (
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/domain/dtos"
	"github.com/muhammadqazi/campus-hq-api/src/internal/core/infrastructure/postgres/entities"
	"gorm.io/gorm"
)

type AccountingRepository interface {
	InsertAccounts(entities.AccountsEntity) (uint, error)
	QueryAccountDetailsByStudentID(uint) ([]dtos.AccountDetails, error)
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

func (r *accountingConnection) QueryAccountDetailsByStudentID(sid uint) ([]dtos.AccountDetails, error) {

	var info []dtos.AccountDetails
	if err := r.conn.Table("accounts_entity as acc").
		Select(`
		acc.account_id, acc.department_fee, acc.scholarship, acc.discount, acc.discount_type, acc.installments,
		acc.total_fee, acc.total_dept, acc.current_dept, acc.approaching_dept,
		pay.date AS payment_date, pay.currency AS payment_currency, pay.installment AS installment_number,
		inv.date AS invoice_date, inv.amount AS invoice_amount, inv.description AS invoice_description,
		inv.installment AS invoice_installment, inv.term
	`).
		Joins(`
		JOIN payments_entity as pay ON pay.account_id = acc.account_id
		JOIN invoices_entity as inv ON inv.account_id = acc.account_id
	`).
		Where("student_id = ?", sid).
		Scan(&info).Error; err != nil {
		return nil, err
	}

	return info, nil
}

func (r *accountingConnection) InsertPayment(payment entities.PaymentsEntity, sid uint) error {

	//account, err := r.QueryAccountDetailsByStudentID(sid)
	//if err != nil {
	//	return err
	//}
	//
	//accountID := account.AccountID
	//payment.AccountID = accountID
	//
	//// Begin a new transaction
	//tx := r.conn.Begin()
	//
	//// Insert the payment
	//if err := tx.Create(&payment).Error; err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//
	//// Update the account total debt
	//account.TotalDept -= payment.Amount
	//if err := tx.Save(&account).Error; err != nil {
	//	tx.Rollback()
	//	return err
	//}

	// Commit the transaction
	//tx.Commit()

	return nil
}
