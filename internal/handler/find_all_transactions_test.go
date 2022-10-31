package handler_test

import (
	"fmt"
	"testing"

	"github.com/kwakubiney/bank-transfer/internal/handler"
	"github.com/stretchr/testify/assert"
)

func TestFindAllTransactionsEndpoint200(t *testing.T) {

	transaction := handler.CreateTestTransaction(t)
	handler.SeedDB(transaction)

	req := handler.MakeTestRequest(t, fmt.Sprintf("/transaction?id=%s", transaction.Credit), nil, "GET")
	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)

	assert.Equal(t, "transactions successfully retreived",
		responseBody["message"])
}

func TestFindAllTransactionsEndpoint404(t *testing.T) {

	transaction := handler.CreateTestTransaction(t)
	handler.SeedDB(transaction)

	req := handler.MakeTestRequest(t, fmt.Sprintf("/transaction?id=%s", "dsfsfsfsa"), nil, "GET")
	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)

	assert.Equal(t, "no transaction found for account specified",
		responseBody["message"])
}
