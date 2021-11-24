package coinbaseclient

import (
	"fmt"
	"os"

	logger "github.com/adamavixio/logger"
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
		logger.HandleError("env var error", err)
	}
	return value
}
