package client

import (
	"fmt"
	"os"
	"strconv"
)

func getEnvVar(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("%s not defined", key)
	}

	return value, nil
}

func FormatMinSize(precision string, amount float64) (string, error) {
	place, dot := 0, false
	for _, char := range precision {
		if !dot && char == '1' {
			break
		}

		if dot && char == '1' {
			place += 1
			break
		}

		if dot && char == '0' {
			place += 1
			continue
		}

		if char == '.' {
			dot = true
		}
	}

	size := strconv.FormatFloat(amount, 'f', place, 64)
	return size, nil
}
