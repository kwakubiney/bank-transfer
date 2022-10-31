package handler_test

import (
	"fmt"
	"testing"

	"github.com/kwakubiney/bank-transfer/internal/handler"
	"github.com/stretchr/testify/assert"
)

func TestFindTransactionsEndpoint200(t *testing.T) {

	transaction := handler.CreateTestTransaction(t)
	handler.SeedDB(transaction)

	req := handler.MakeTestRequest(t, fmt.Sprintf("/transaction/filter?credit=%s", transaction.Credit), nil, "GET")
	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)

	assert.Equal(t, "transactions successfully retreived",
		responseBody["message"])
}

func TestFindTransactionsEndpoint404(t *testing.T) {

	transaction := handler.CreateTestTransaction(t)
	handler.SeedDB(transaction)

	req := handler.MakeTestRequest(t, fmt.Sprintf("/transaction/filter?credit=%s", "dsfsfsfsa"), nil, "GET")
	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)
	assert.Equal(t, "no transaction found for filters specified",
		responseBody["message"])
}
