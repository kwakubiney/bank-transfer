package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"log"

	"github.com/kwakubiney/bank-transfer/internal/domain/model"
	"github.com/kwakubiney/bank-transfer/internal/domain/repository"
)

type FindAllTransferRequest struct{
	ID string `form:"id" `
}

func (h *Handler) FindAllTransactions(c *gin.Context) {
	var transaction model.Transaction
	var FindAllTransferRequest FindAllTransferRequest
	err := c.BindQuery(&FindAllTransferRequest)
	//TODO: Validate ID field such that no other param can be passed
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check API documentation",
		})
		return
	}
	newTransaction, err := h.TransactionRepo.FindAllTransactions(FindAllTransferRequest.ID, &transaction)
	if err != nil {
		log.Println(err)
		if err == repository.ErrTransactionNotFound{
			c.JSON(http.StatusNotFound, gin.H{
				"message": "no transaction found for account specified",
			})
			return
		}else{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "service error",
			})
			return
		}	
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "transactions successfully retreived",
		"transactions": newTransaction,
	})
}



