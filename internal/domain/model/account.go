package model

import (
	"errors"
	"time"
)

var (
	ErrInsufficientBalance = errors.New("insufficient balance in account")
)

type Account struct {
	ID            string    `json:"id" gorm:"default:gen_random_uuid()"`
	Name          string    `json:"name"`
	AccountNumber string    `json:"account_number"`
	Balance       Money     `json:"balance"`
	CreatedAt     time.Time `json:"created_at" sql:"type:timestamp without time zone"`
}

func NewAccount(id string, name string, accountNumber string, balance Money) *Account {
	return &Account{
		ID:            id,
		Name:          name,
		AccountNumber: accountNumber,
		Balance:       balance,
		CreatedAt:     time.Now(),
	}
}

func (a *Account) Deposit(amount Money) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount Money) error {
	if a.Balance < amount {
		return ErrInsufficientBalance
	}

	a.Balance -= amount

	return nil
}

func NewAccountBalance(balance Money) Account {
	return Account{Balance: balance}
}
