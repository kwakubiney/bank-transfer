package handler

import (
	"log"

	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/kwakubiney/bank-transfer/internal/domain/model"
	"github.com/kwakubiney/bank-transfer/internal/domain/repository"
)

func (h *Handler) FindTransaction(c *gin.Context) {
	var transaction model.Transaction
	queryParams := c.Request.URL.Query()

	//TODO: Validate params such that no other param can be passed
	//Currently expecting "credit=xxx" or "debit=xxxx"
	newTransaction, err := h.TransactionRepo.FindTransactions(queryParams, &transaction)
	if err != nil {
		if err != nil {
			log.Println(err)
			if err == repository.ErrTransactionNotFound{
				c.JSON(http.StatusNotFound, gin.H{
					"message": "no transaction found for filters specified",
				})
				return
			}else{
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "service error",
				})
				return
			}	
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "transactions successfully retreived",
		"transactions": newTransaction,
	})








	
	}