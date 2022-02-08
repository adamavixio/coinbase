package client

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

func TestAccountByID(t *testing.T) {
	id, err := getEnvVar("COINBASE_ACCOUNT_ID_USD")
	if err != nil {
		t.Error(err)
	}

	res, err := AccountByID(&AccountByIDConfig{id})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
