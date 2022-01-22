package coinbaseclient

import (
	"fmt"
	"testing"
)

func TestAccounts(t *testing.T) {
	res, err := Accounts()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestAccountUSD(t *testing.T) {
	res, err := AccountByCurrency(USD)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}

func TestAccountByID(t *testing.T) {
	id, err := getEnvVar("COINBASE_ACCOUNT_ID_USD")
	if err != nil {
		t.Error(err)
	}

	config := &AccountByIDConfig{
		ID: id,
	}

	res, err := AccountByID(config)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
