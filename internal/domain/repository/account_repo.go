package repository

import (
	//"gorm.io/gorm"
	"errors"
	"log"

	"github.com/kwakubiney/bank-transfer/internal/domain/model"
	"gorm.io/gorm"
	"time"
)

var (
	ErrAccountNotFound            = errors.New("account not found")
	ErrAccountOriginNotFound      = errors.New("account origin not found")
	ErrAccountDestinationNotFound = errors.New("destination not found")
	ErrAccountCannotBeCreated     = errors.New("account cannot be created")
)

type AccountRepository struct {
	DB *gorm.DB
}


func (a *AccountRepository) WithTrx(trxHandle *gorm.DB) *AccountRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return &AccountRepository{}
	}
	a.DB = trxHandle
	return a
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		db,
	}
}

func (a *AccountRepository) CreateAccount(account *model.Account) error {
	if a.DB.Create(&account).Error != nil {
		return ErrAccountCannotBeCreated
	}
	return nil
}

func (a *AccountRepository) DepositToAccount(account *model.Account, amount int64) (int64, int64, error) {
	db := a.DB.Model(model.Account{}).Find(&account)
	if db.RowsAffected == 0 {
		return 0, 0, ErrAccountOriginNotFound
	}

	oldBalance := account.Balance
	account.Deposit(amount)

	account.LastModified = time.Now()
	err := db.Model(&account).Update("balance", account.Balance).Error
	if err != nil {
		log.Println(db.Error)
		return 0, 0, err
	}
	err = db.Model(&account).Update("last_modified", account.LastModified).Error
	if err != nil {
		log.Println(db.Error)
		return 0, 0, err
	}
	newBalance := account.Balance
	return oldBalance, newBalance, nil
}

func (a *AccountRepository) WithdrawFromAccount(account *model.Account, amount int64) (int64, int64, error) {
	db := a.DB.Model(model.Account{}).Find(&account)
	if db.RowsAffected == 0 {
		return 0, 0, ErrAccountOriginNotFound
	}

	oldBalance := account.Balance
	err := account.Withdraw(amount)
	if err != nil {
		return 0, 0, err
	}

	//TODO: Bundle both queries into one
	account.LastModified = time.Now()
	err = db.Model(&account).Update("balance", account.Balance).Error
	if err != nil {
		log.Println(db.Error)
		return 0, 0, err
	}
	err = db.Model(&account).Update("last_modified", account.LastModified).Error
	if err != nil {
		log.Println(db.Error)
		return 0, 0, err
	}
	newBalance := account.Balance
	return oldBalance, newBalance, nil
}

func (a *AccountRepository) FindAccountByID(originID string, destinationID string) (*model.Account, *model.Account, error) {
	var originAccount *model.Account
	var destinationAccount *model.Account
	db := a.DB.Where("id = ?", originID).Find(&originAccount)
	if db.RowsAffected == 0 {
		return nil, nil, ErrAccountOriginNotFound
	}
	db = a.DB.Where("id = ?", destinationID).Find(&destinationAccount)
	if db.RowsAffected == 0 {
		return nil, nil, ErrAccountDestinationNotFound
	}
	return originAccount, destinationAccount, nil
}

func (a *AccountRepository) Transfer(origin model.Account, destination model.Account, amount int64) error {
	err := origin.Withdraw(amount)
	if err != nil {
		return err
	}

	err = a.DB.Model(&origin).Update("balance", origin.Balance).Error
	if err != nil {
		log.Println(err)
		return err
	}

	destination.Deposit(amount)
	err = a.DB.Model(&destination).Update("balance", destination.Balance).Error
	if err != nil {
		log.Println(err)
		return err
	}

	//TODO: Update last modified on both accounts
	return nil
}


