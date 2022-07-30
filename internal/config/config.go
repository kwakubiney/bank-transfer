package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadNormalConfig() error {
	pathToEnvFile := "./.env"
	if _, err := os.Stat(pathToEnvFile); err == nil {
		return godotenv.Load(pathToEnvFile)
	}

	return nil
}
