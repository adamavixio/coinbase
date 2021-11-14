package coinbaseapi

import (
	"fmt"
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
	value := os.Getenv(key)
	if value == "" {
		err := fmt.Errorf("%s not defined", key)
		handleError("env var error", err)
	}
	return value
}

func handleError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
