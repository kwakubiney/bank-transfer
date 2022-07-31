package handler

import "github.com/kwakubiney/bank-transfer/internal/domain/repository"

type Handler struct {
	AccountRepo repository.AccountRepository
}

func NewHandler(accountRepo *repository.AccountRepository) *Handler {
	return &Handler{
		AccountRepo: *accountRepo,
	}
}
