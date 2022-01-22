package coinbaseclient

import (
	"encoding/json"
	"fmt"
)

//
// Implementation
//

type Account struct {
	ID             string `json:"id"`
	Currency       string `json:"currency"`
	Balance        string `json:"balance"`
	Available      string `json:"available"`
	Hold           string `json:"hold"`
	ProfileID      string `json:"profile_id"`
	TradingEnabled bool   `json:"trading_enabled"`
}

type AccountByIDConfig struct {
	ID string `json:"id"`
}

func (accountByIDConfig *AccountByIDConfig) isValid() error {
	if accountByIDConfig.ID == "" {
		return fmt.Errorf("invalid value for field ID: %s", accountByIDConfig.ID)
	}

	return nil
}

//
// API
//

func Accounts() ([]Account, error) {
	request := RequestConfig{
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

func AccountByCurrency(currency string) (*Account, error) {
	accounts, err := Accounts()
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		if account.Currency == currency {
			return &account, nil
		}
	}

	return nil, fmt.Errorf("unable to find account ID for USD")
}

func AccountByID(config *AccountByIDConfig) (*Account, error) {
	err := config.isValid()
	if err != nil {
		return nil, err
	}

	request := RequestConfig{
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
