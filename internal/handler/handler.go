package handler

import "github.com/kwakubiney/bank-transfer/internal/domain/repository"

type Handler struct {
	AccountRepo repository.AccountRepository
	TransactionRepo repository.TransactionRepository
}

func NewHandler(accountRepo *repository.AccountRepository, transactionRepo *repository.TransactionRepository) *Handler {
	return &Handler{
		AccountRepo: *accountRepo,
		TransactionRepo: *transactionRepo,
	}
}

