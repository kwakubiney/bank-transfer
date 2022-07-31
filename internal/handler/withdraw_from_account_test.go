package handler_test

import (
	"testing"
	"github.com/kwakubiney/bank-transfer/internal/handler"
	"github.com/stretchr/testify/assert"
)

func TestWithdrawFromAccountEndpoint200(t *testing.T) {

	account := handler.CreateTestAccount(t)
	handler.SeedDB(account)

	req := handler.MakeTestRequest(t, "/withdraw", map[string]interface{}{
		"id": account.ID,
		"amount": 100,
	}, "POST")

	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)

// Type cast since json.Unmarshal converts JSON numbers to floats.
	assert.Equal(t, map[string]interface {}{"amount":float64(100), "message":"amount successfully withdrawn", "new_balance":float64(100), "old_balance":float64(200)}, responseBody)
}

func TestWithdrawFromAccountEndpoint400(t *testing.T) {

	account := handler.CreateTestAccount(t)
	handler.SeedDB(account)

	req := handler.MakeTestRequest(t, "/withdraw", map[string]interface{}{
		"id": account.ID,
		"amount": 2000,
	}, "POST")

	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)
	assert.Equal(t, map[string]interface {}{"message":"could not withdraw amount due to insufficient balance"}, responseBody)
}

func TestWithdrawFromAccountEndpoint404(t *testing.T) {

	account := handler.CreateTestAccount(t)
	handler.SeedDB(account)

	req := handler.MakeTestRequest(t, "/withdraw", map[string]interface{}{
		"id": "nfsjknfskjnfksanf",
		"amount": 2000,
	}, "POST")

	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)
	assert.Equal(t, map[string]interface {}{"message":"account origin specified was not found"}, responseBody)
}
