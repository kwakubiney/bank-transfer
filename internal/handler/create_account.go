package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/bank-transfer/internal/domain/model"
	"github.com/kwakubiney/bank-transfer/internal/domain/repository"
	"log"
	"net/http"
	"time"
)

type CreateAccountRequest struct {
	ID        string    `json:"id" gorm:"default:gen_random_uuid()"`
	Name      string    `json:"name" binding:"required"`
	Balance   int64     `json:"balance" binding:"required"`
	CreatedAt time.Time `json:"created_at" sql:"type:timestamp without time zone"`
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var account model.Account
	var createAccountRequest CreateAccountRequest
	err := c.BindJSON(&createAccountRequest)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check API documentation",
		})
		return
	}

	account.Balance = createAccountRequest.Balance
	account.Name = createAccountRequest.Name
	account.CreatedAt = time.Now()
	account.LastModified = time.Now()

	err = h.AccountRepo.CreateAccount(&account)
	if err == repository.ErrAccountCannotBeCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not create account",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "account successfully created",
		"account": account,
	})
}
