package coinbaseclient

import (
	"fmt"
	"os"
)

const (
	get    = "GET"
	post   = "POST"
	url    = "https://api.pro.coinbase.com"
	socket = "wss://ws-feed.pro.coinbase.com"
)

func getEnvVar(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("%s not defined", key)
	}

	return value, nil
}
