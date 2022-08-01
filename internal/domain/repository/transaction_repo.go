package repository

import (
	"github.com/kwakubiney/bank-transfer/internal/domain/model"
	"gorm.io/gorm"
	"errors"
	"fmt"
	"net/url"
)

var (
	ErrTransactionNotFound            = errors.New("transaction not found")
	ErrTransactionCannotBeCreated     = errors.New("transaction cannot be created")
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db,
	}
}

func (t *TransactionRepository) CreateTransaction(transaction *model.Transaction) error{
	if t.db.Create(&transaction).Error != nil {
	return ErrTransactionCannotBeCreated
}
	return nil
}

func (t *TransactionRepository) FindAllTransactions(accountID string, transaction *model.Transaction) (*model.Transaction, error){
	db := t.db.Where(fmt.Sprintf("credit = '%s'", accountID)).Or(map[string]interface{}{"debit": accountID}).Find(&transaction) 
	if db.RowsAffected == 0 {
			return nil, ErrTransactionNotFound
		}
	return transaction, nil
}

func (t *TransactionRepository) FindTransactions(queryString url.Values , transaction *model.Transaction) (*model.Transaction, error){
	newMap := make(map[string]interface{})
	for k, v := range queryString{
        newMap[k] = v[0]
    }
	db := t.db.Where(newMap).Find(&transaction)
	if db.RowsAffected == 0 {
			return nil, ErrTransactionNotFound
		}
	return transaction, nil
}
