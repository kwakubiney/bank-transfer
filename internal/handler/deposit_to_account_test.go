package handler_test

import (
	"github.com/kwakubiney/bank-transfer/internal/handler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDepositToAccountEndpoint200(t *testing.T) {

	account := handler.CreateTestAccount(t)
	handler.SeedDB(account)

	req := handler.MakeTestRequest(t, "/deposit", map[string]interface{}{
		"id":     account.ID,
		"amount": 100,
	}, "POST")

	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)

	// Type cast since json.Unmarshal converts JSON numbers to floats.
	assert.Equal(t, map[string]interface{}{"amount": float64(100), "message": "amount successfully deposited", "new_balance": float64(300), "old_balance": float64(200)}, responseBody)
}

func TestDepositToAccountEndpoint404(t *testing.T) {

	account := handler.CreateTestAccount(t)
	handler.SeedDB(account)

	req := handler.MakeTestRequest(t, "/deposit", map[string]interface{}{
		"id":     "isjhfsjakfnkasnfkas",
		"amount": 2000,
	}, "POST")

	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)
	assert.Equal(t, map[string]interface{}{"message": "account origin specified was not found"}, responseBody)
}
