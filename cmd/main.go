package main

import (
	"github.com/kwakubiney/bank-transfer/internal/config"
	"github.com/kwakubiney/bank-transfer/internal/postgres"
	"github.com/kwakubiney/bank-transfer/internal/server"
	"log"
)

func main() {
	err := config.LoadNormalConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, err = postgres.Init()
	if err != nil {
		log.Fatal(err)
	}
	server := server.New()
	server.Start()
}
