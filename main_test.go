package main

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestBotToken(t *testing.T) {
	configErr := godotenv.Load("configs/config.env")

	if configErr != nil {
		t.Fatalf("Error loading .env file")
	}

	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	if telegramBotToken == "" {
		t.Fatalf("Token variable not set")
	}
}
