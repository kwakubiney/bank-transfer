package repository

import (
	//"gorm.io/gorm"
	"errors"
	"log"

	"github.com/kwakubiney/bank-transfer/internal/domain/model"
	"gorm.io/gorm"
)

var (
	ErrAccountNotFound       = errors.New("account not found")
	ErrAccountOriginNotFound = errors.New("account origin not found")
	ErrAccountCannotBeCreated = errors.New("account cannot be created")
	
)
type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository (db *gorm.DB) *AccountRepository{
	return &AccountRepository{
		db,
	}
}

func (a *AccountRepository) CreateAccount(account *model.Account) error {
	if a.db.Create(&account).Error != nil {
		return ErrAccountCannotBeCreated
	}
	return nil
}

func (a *AccountRepository) DepositToAccount(account *model.Account, amount model.Money)(model.Money, model.Money, error) {
	db :=  a.db.Model(model.Account{}).Where("id = ?", account.ID).Find(&account)
	if db.RowsAffected == 0 {
		return 0, 0, ErrAccountOriginNotFound
	}

	oldBalance := account.Balance
	account.Deposit(amount)

	err := db.Updates(&account).Error
	if err != nil {
		log.Println(db.Error)
		return 0, 0, err
	}
	newBalance := account.Balance
	return oldBalance, newBalance, nil
}

func (a *AccountRepository) WithdrawFromAccount(account *model.Account, amount model.Money) (model.Money, model.Money, error) {
	db :=  a.db.Model(model.Account{}).Where("id = ?", account.ID).Find(&account)
	if db.RowsAffected == 0 {
		return 0, 0, ErrAccountOriginNotFound
	}

	oldBalance := account.Balance
	err := account.Withdraw(amount)
	if err != nil {
		return 0, 0, err
	}

	err = db.Updates(&account).Error
	if err != nil {
		log.Println(db.Error)
		return 0, 0, err
	}
	newBalance := account.Balance
	return oldBalance, newBalance, nil
}


func (a *AccountRepository) UpdateBalance (account model.Account, action string){

}

func (a *AccountRepository) FindAll (account model.Account){
	//
}

func (a *AccountRepository) FindAccountByID (account model.Account, id string){
	//
}

func (a *AccountRepository) FindBalanceByID (account model.Account, id string){
	
}




