package repository

import (
	"errors"
	"fmt"
	"log"
	"github.com/kwakubiney/bank-transfer/internal/domain/model"
	"gorm.io/gorm"
	"net/url"
)

var (
	ErrTransactionNotFound        = errors.New("transaction not found")
	ErrTransactionCannotBeCreated = errors.New("transaction cannot be created")
)

type TransactionRepository struct {
	DB *gorm.DB
}


func (t *TransactionRepository) WithTrx(trxHandle *gorm.DB) *TransactionRepository{
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return &TransactionRepository{}
	}
	t.DB = trxHandle
	return t
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db,
	}
}

func (t *TransactionRepository) CreateTransaction(transaction *model.Transaction) error {
	if t.DB.Create(&transaction).Error != nil {
		return ErrTransactionCannotBeCreated
	}
	return nil
}

func (t *TransactionRepository) FindAllTransactions(accountID string, transaction *model.Transaction) (*model.Transaction, error) {
	db := t.DB.Where(fmt.Sprintf("credit = '%s'", accountID)).Or(map[string]interface{}{"debit": accountID}).Find(&transaction)
	if db.RowsAffected == 0 {
		return nil, ErrTransactionNotFound
	}
	return transaction, nil
}

func (t *TransactionRepository) FindTransactions(queryString url.Values, transaction *model.Transaction) (*model.Transaction, error) {
	newMap := make(map[string]interface{})
	for k, v := range queryString {
		newMap[k] = v[0]
	}
	db := t.DB.Where(newMap).Find(&transaction)
	if db.RowsAffected == 0 {
		return nil, ErrTransactionNotFound
	}
	return transaction, nil
}
