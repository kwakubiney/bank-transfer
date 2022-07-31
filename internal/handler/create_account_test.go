package handler_test

import (
	"log"
	"os"

	"testing"
	//"github.com/jaswdr/faker"
	"github.com/gin-gonic/gin"
	"github.com/kwakubiney/bank-transfer/internal/config"
	"github.com/kwakubiney/bank-transfer/internal/domain/repository"
	"github.com/kwakubiney/bank-transfer/internal/handler"
	"github.com/kwakubiney/bank-transfer/internal/postgres"
	"github.com/kwakubiney/bank-transfer/internal/server"
	"github.com/stretchr/testify/assert"
)

var engine *gin.Engine

func TestMain(m *testing.M) {
	err := config.LoadTestConfig()
	assert.NoError(&testing.T{}, err)
	db, err := postgres.Init()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewAccountRepository(db)
	handler := handler.NewHandler(repo)
	server := server.New(handler)
	engine = server.SetupRoutes()
	os.Exit(m.Run())
}

func TestCreateAccountEndpoint200(t *testing.T) {
	req := handler.MakeTestRequest(t, "/createAccount", map[string]interface{}{
		"name":    "kwamz",
		"balance": 400,
	}, "POST")

	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)
	assert.Equal(t, "account successfully created", responseBody["message"])
}

func TestCreateAccountEndpoint400(t *testing.T) {
	req := handler.MakeTestRequest(t, "/createAccount", map[string]interface{}{
		"balance": 400,
	}, "POST")

	response := handler.BootstrapServer(req, engine)
	responseBody := handler.DecodeResponse(t, response)
	assert.Equal(t, "could not parse request. check API documentation", responseBody["message"])
}
