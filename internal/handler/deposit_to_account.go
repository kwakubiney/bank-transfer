package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/bank-transfer/internal/domain/model"
	"github.com/kwakubiney/bank-transfer/internal/domain/repository"
)

type DepositToAccountRequest struct {
	ID     string `json:"id" gorm:"default:gen_random_uuid()" binding:"required"`
	Amount int64  `json:"amount" binding:"required,gt=0"`
}

func (h *Handler) DepositToAccount(c *gin.Context) {
	//TODO: Make changes to transaction table
	var account model.Account
	var DepositToAccountRequest DepositToAccountRequest
	err := c.BindJSON(&DepositToAccountRequest)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check API documentation",
		})
		return
	}

	account.ID = DepositToAccountRequest.ID
	account.LastModified = time.Now()
	amount := DepositToAccountRequest.Amount
	oldBalance, newBalance, err := h.AccountRepo.DepositToAccount(&account, amount)
	if err != nil {
		if err == repository.ErrAccountOriginNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "account origin specified was not found",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "could not deposit amount",
			})

			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "amount successfully deposited",
		"old_balance": oldBalance,
		"new_balance": newBalance,
		"amount":      amount,
	})
}
