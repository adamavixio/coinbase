package coinbaseclient

import (
	"fmt"
	"os"
)

const (
	url    = "https://api.exchange.coinbase.com"
	socket = "wss://ws-feed.pro.coinbase.com"

	get    = "GET"
	post   = "POST"
	delete = "DELETE"

	USD = "USD"
)

func getEnvVar(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("%s not defined", key)
	}

	return value, nil
}
