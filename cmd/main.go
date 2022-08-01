package main

import (
	"log"

	"github.com/kwakubiney/bank-transfer/internal/config"
	"github.com/kwakubiney/bank-transfer/internal/domain/repository"
	"github.com/kwakubiney/bank-transfer/internal/handler"
	"github.com/kwakubiney/bank-transfer/internal/postgres"
	"github.com/kwakubiney/bank-transfer/internal/server"
)

func main() {
	err := config.LoadNormalConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.Init()
	if err != nil {
		log.Fatal(err)
	}

	accountRepo := repository.NewAccountRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	handlers := handler.NewHandler(accountRepo, transactionRepo)
	server := server.New(handlers)
	server.Start()
}
