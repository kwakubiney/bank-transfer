package repository

import (
	//"gorm.io/gorm"
	"errors"

	"github.com/kwakubiney/bank-transfer/internal/domain/model"
	"gorm.io/gorm"
)

var (
	ErrAccountNotFound       = errors.New("account not found")
	ErrAccountOriginNotFound = errors.New("account origin not found")
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository (db *gorm.DB) *AccountRepository{
	return &AccountRepository{
		db,
	}
}

func (a *AccountRepository) CreateAccount(account model.Account) error {
	return a.db.Create(&account).Error
}