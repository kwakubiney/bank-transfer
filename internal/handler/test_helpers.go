package handler

import (
	"bytes"
	"encoding/json"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"

	"github.com/kwakubiney/bank-transfer/internal/domain/model"
	"github.com/kwakubiney/bank-transfer/internal/postgres"
)

var dbConnPool *gorm.DB

func NewUUID() string {
	uid, _ := uuid.NewV4()
	return uid.String()
}

func MakeRequest(route string, port string, requestBody interface{}, method string) (*http.Response, error) {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, route, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func SeedDB(r ...interface{}) *gorm.DB {
	if dbConnPool == nil {
		db, err := postgres.Init()
		if err != nil {
			log.Fatal(err)
		}
		dbConnPool = db
	}
	tx := dbConnPool.Begin()
	for _, m := range r {
		if err := tx.Create(m).Error; err != nil {
			tx.Rollback()
			log.Fatalf("[data insert failed] %v ", err)
		}
	}
	tx.Commit()
	return dbConnPool
}

func CreateTestAccount(t *testing.T) *model.Account {
	f := faker.New()

	testAccount := model.Account{
		ID:        NewUUID(),
		Name:      f.Internet().User(),
		Balance:   200,
		CreatedAt: time.Now(),
	}

	return &testAccount
}

func CreateTestTransaction(t *testing.T) *model.Transaction {

	testTransaction := model.Transaction{
		ID:        NewUUID(),
		Credit:     NewUUID(),
		Debit:   NewUUID(),
		Amount: faker.New().Int64(),
		CreatedAt: time.Now(),
	}
	return &testTransaction
}

func BootstrapServer(req *http.Request, routeHandlers *gin.Engine) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	routeHandlers.ServeHTTP(responseRecorder, req)
	return responseRecorder
}

func MakeTestRequest(t *testing.T, route string, body interface{}, method string) *http.Request {
	if body == nil {
		req, err := http.NewRequest(method, route, nil)
		assert.NoError(t, err)
		return req
	}

	reqBody, err := json.Marshal(body)
	assert.NoError(t, err)
	req, err := http.NewRequest(method, route, bytes.NewReader(reqBody))
	assert.NoError(t, err)
	return req
}

func DecodeResponse(t *testing.T, response *httptest.ResponseRecorder) map[string]interface{} {
	var responseBody map[string]interface{}
	assert.NoError(t, json.Unmarshal(response.Body.Bytes(), &responseBody))
	return responseBody
}
