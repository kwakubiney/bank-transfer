package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/bank-transfer/internal/domain/model"
	"github.com/kwakubiney/bank-transfer/internal/domain/repository"
	"gorm.io/gorm"
)

type WithdrawFromAccountRequest struct {
	ID     string `json:"id" gorm:"default:gen_random_uuid()" binding:"required"`
	Amount int64  `json:"amount" binding:"required,gt=0"`
}

func (h *Handler) WithdrawFromAccount(c *gin.Context) {
	var transaction model.Transaction
	var account model.Account
	var WithdrawFromAccountRequest WithdrawFromAccountRequest
	err := c.BindJSON(&WithdrawFromAccountRequest)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check API documentation",
		})
		return
	}

	account.ID = WithdrawFromAccountRequest.ID

	amount := WithdrawFromAccountRequest.Amount
	tx := c.MustGet("db_trx").(*gorm.DB)
	oldBalance, newBalance, err := h.AccountRepo.WithTrx(tx).WithdrawFromAccount(&account, amount)
	if err != nil {
		if err == model.ErrInsufficientBalance {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "could not withdraw amount due to insufficient balance",
			})
			return
		}
		if err == repository.ErrAccountOriginNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "account origin specified was not found",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "could not withdraw amount",
			})

			return
		}
	}

	transaction.Amount = amount
	transaction.CreatedAt = time.Now()
	transaction.Debit = WithdrawFromAccountRequest.ID

	err = h.TransactionRepo.WithTrx(tx).CreateTransaction(
		&transaction,
	)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not deposit amount",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "amount successfully withdrawn",
		"old_balance": oldBalance,
		"new_balance": newBalance,
		"amount":      amount,
	})
}
