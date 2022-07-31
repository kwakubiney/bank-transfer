package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/bank-transfer/internal/domain/model"
	"github.com/kwakubiney/bank-transfer/internal/domain/repository"
)

type TransferRequest struct {
	OriginAccountID      string `json:"origin" binding:"required"`
	DestinationAccountID string `json:"destination" binding:"required"`
	Amount               int64  `json:"amount" binding:"required,gt=0"`
}

func (h *Handler) TransferToAccount(c *gin.Context) {
	var transferRequest TransferRequest
	err := c.BindJSON(&transferRequest)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check API documentation",
		})
		return
	}

	originAccount, destinationAccount, err := h.AccountRepo.FindAccountByID(transferRequest.OriginAccountID, transferRequest.DestinationAccountID)
	if err != nil {
		if err == repository.ErrAccountOriginNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "account origin specified was not found",
			})
			return
		}

		if err == repository.ErrAccountDestinationNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "account destination specified was not found",
			})
			return
		} else {
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	err = h.AccountRepo.UpdateBalanceAfterTransfer(*originAccount, *destinationAccount, transferRequest.Amount)
	if err != nil {
		log.Println(err)
		if err == model.ErrInsufficientBalance {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "could not withdraw amount due to insufficient balance",
			})
			return
		} else {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "settlement failed on both accounts"},
			)
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"account_origin_id": transferRequest.OriginAccountID,
		"account_destination_id": transferRequest.DestinationAccountID,
		"amount": transferRequest.Amount,
	})

}
