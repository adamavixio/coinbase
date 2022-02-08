package client

import (
	"fmt"
	"strconv"
)

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

func BalanceByCurrency(currency string) (float64, error) {
	account, err := AccountByCurrency(currency)
	if err != nil {
		return 0, err
	}

	balance, err := strconv.ParseFloat(account.Balance, 64)
	if err != nil {
		return 0, err
	}

	return balance, nil
}
