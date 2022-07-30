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
	Balance       Money     `json:"balance"`
	CreatedAt     time.Time `json:"created_at" sql:"type:timestamp without time zone"`
	LastModified  time.Time `json:"last_modified" sql:"type:timestamp without time zone"`
}

func NewAccount(id string, name string, accountNumber string, balance Money) *Account {
	return &Account{
		ID:            id,
		Name:          name,
		Balance:       balance,
		CreatedAt:     time.Now(),
		LastModified:  time.Now(),	
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