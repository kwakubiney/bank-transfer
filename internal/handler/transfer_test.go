package handler_test

import (
	"github.com/kwakubiney/bank-transfer/internal/handler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransferToAccountEndpoint200(t *testing.T) {

	accountOrigin := handler.CreateTestAccount(t)
	accountDestination := handler.CreateTestAccount(t)
	handler.SeedDB(accountOrigin, accountDestination)

	req := handler.MakeTestRequest(t, "/transfer", map[string]interface{}{
		"origin":      accountOrigin.ID,
		"destination": accountDestination.ID,
		"amount":      100,
	}, "POST")

	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)

	// Type cast since json.Unmarshal converts JSON numbers to floats.
	assert.Equal(t, map[string]interface{}{"account_destination_id": accountDestination.ID, "account_origin_id": accountOrigin.ID, "amount": float64(100)}, responseBody)
}

func TestTransferToAccountEndpoint400(t *testing.T) {
	accountOrigin := handler.CreateTestAccount(t)
	accountDestination := handler.CreateTestAccount(t)
	handler.SeedDB(accountOrigin, accountDestination)
	req := handler.MakeTestRequest(t, "/transfer", map[string]interface{}{
		"origin":      accountOrigin.ID,
		"destination": accountDestination.ID,
		"amount":      300,
	}, "POST")

	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)
	assert.Equal(t, map[string]interface{}{"message": "could not withdraw amount due to insufficient balance"}, responseBody)
}

func TestTransferToAccountEndpointNegativeAmount400(t *testing.T) {
	accountOrigin := handler.CreateTestAccount(t)
	accountDestination := handler.CreateTestAccount(t)
	handler.SeedDB(accountOrigin, accountDestination)
	req := handler.MakeTestRequest(t, "/transfer", map[string]interface{}{
		"origin":      accountOrigin.ID,
		"destination": accountDestination.ID,
		"amount":      -40,
	}, "POST")

	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)
	assert.Equal(t, map[string]interface{}{"message": "could not parse request. check API documentation"}, responseBody)
}

func TestTransferToAccountEndpoint404(t *testing.T) {
	accountOrigin := handler.CreateTestAccount(t)
	accountDestination := handler.CreateTestAccount(t)
	handler.SeedDB(accountOrigin, accountDestination)
	req := handler.MakeTestRequest(t, "/transfer", map[string]interface{}{
		"origin":      "hjdhsakjfcnsjfs",
		"destination": accountDestination.ID,
		"amount":      10}, "POST")
	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)
	assert.Equal(t, map[string]interface{}{"message": "account origin specified was not found"}, responseBody)
}
