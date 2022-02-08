package client

import "fmt"

type AccountByIDConfig struct {
	ID string `json:"id"`
}

func (config *AccountByIDConfig) validate() error {
	if config.ID == "" {
		return fmt.Errorf("invalid value for field ID: %s", config.ID)
	}

	return nil
}
