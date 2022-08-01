package repository

import (
	"github.com/kwakubiney/bank-transfer/internal/domain/model"
	"gorm.io/gorm"
	"errors"
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

func (t *TransactionRepository) FilterTransactionByUser(transaction *model.Transaction) error{
	if t.db.Create(&transaction).Error != nil {
	return ErrTransactionCannotBeCreated
}
	return nil
}
