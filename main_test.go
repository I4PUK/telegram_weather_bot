package main

import (
	"github.com/joho/godotenv"
	"testing"
)

func TestReadEnvFile(t *testing.T) {
	configErr := godotenv.Load("configs/config.env")

	if configErr != nil {
		t.Fatalf("Error loading .env file")
	}
}
