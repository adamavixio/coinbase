package client

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Available      string `json:"available"`
	Balance        string `json:"balance"`
	Currency       string `json:"currency"`
	Hold           string `json:"hold"`
	ID             string `json:"id"`
	ProfileID      string `json:"profile_id"`
	TradingEnabled bool   `json:"trading_enabled"`
}

func Accounts() ([]Account, error) {
	request := requestConfig{
		Method: get,
		Path:   "/accounts",
	}

	data, err := authRequest(request)
	if err != nil {
		return nil, err
	}

	account := []Account{}

	err = json.Unmarshal(data, &account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func AccountByID(config *AccountByIDConfig) (*Account, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	request := requestConfig{
		Method: get,
		Path:   fmt.Sprintf("/accounts/%s", config.ID),
	}

	data, err := authRequest(request)
	if err != nil {
		return nil, err
	}

	account := &Account{}

	err = json.Unmarshal(data, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
