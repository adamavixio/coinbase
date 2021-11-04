package coinbaseapi

import (
	"errors"
	"log"
	"os"
)

const (
	get    = "GET"
	post   = "POST"
	url    = "https://api.pro.coinbase.com"
	socket = "wss://ws-feed.pro.coinbase.com"
)

func getEnvVar(key string) string {
	value := os.Getenv("COINBASE_KEY")
	if value == "" {
		handleError("env var error", errors.New("COINBASE_KEY not defined"))
	}
	return value
}

func handleError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
