package model

import (
	"errors"
	"time"
)

var (
	ErrInsufficientBalance = errors.New("insufficient balance in account")
)

type Account struct {
	ID           string    `json:"id" gorm:"default:gen_random_uuid()"`
	Name         string    `json:"name"`
	Balance      int64     `json:"balance" sql:"type:decimal(20,8);"`
	CreatedAt    time.Time `json:"created_at" sql:"type:timestamp without time zone"`
	LastModified time.Time `json:"last_modified" sql:"type:timestamp without time zone"`
}

func NewAccount(id string, name string, accountNumber string, balance int64) *Account {
	return &Account{
		ID:           id,
		Name:         name,
		Balance:      balance,
		CreatedAt:    time.Now(),
		LastModified: time.Now(),
	}
}

func (a *Account) Deposit(amount int64) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount int64) error {

	if a.Balance < amount {
		return ErrInsufficientBalance
	}

	a.Balance -= amount

	return nil
}
