package api

import (
	"fmt"
	"time"

	"github.com/adamavixio/coinbase-service/client"
)

func Buy(productID string, amount float64) error {
	price, err := client.Price(productID)
	if err != nil {
		return err
	}

	minSize, err := client.MinSize(productID)
	if err != nil {
		return err
	}

	size, err := client.FormatMinSize(minSize, amount/price)
	if err != nil {
		return err
	}

	_, err = client.Buy(productID, size, "market")
	return err
}

func BuyAll(productID string) error {
	price, err := client.Price(productID)
	if err != nil {
		return err
	}

	minSize, err := client.MinSize(productID)
	if err != nil {
		return err
	}

	balance, err := client.BalanceByCurrency("USD")
	if err != nil {
		return err
	}

	size, err := client.FormatMinSize(minSize, price/balance)
	if err != nil {
		return err
	}

	_, err = client.Buy(productID, size, "market")
	return err
}

func Sell(productID string, amount float64) error {
	balance, err := client.BalanceByCurrency("BTC")
	if err != nil {
		return err
	}

	if amount > balance {
		return fmt.Errorf("%f greater than %f for %s", amount, balance, productID)
	}

	minSize, err := client.MinSize(productID)
	if err != nil {
		return err
	}

	size, err := client.FormatMinSize(minSize, amount)
	if err != nil {
		return err
	}

	_, err = client.Sell(productID, size, "market")
	return err
}

func SellAll(productID string) error {
	balance, err := client.BalanceByCurrency("BTC")
	if err != nil {
		return err
	}

	minSize, err := client.MinSize(productID)
	if err != nil {
		return err
	}

	size, err := client.FormatMinSize(minSize, balance)
	if err != nil {
		return err
	}

	_, err = client.Sell(productID, size, "market")
	return err
}

func Spread() error {
	products, err := client.USDProductIDs()
	if err != nil {
		return err
	}

	balance, err := client.BalanceByCurrency("USD")
	if err != nil {
		return err
	}

	amount := balance / float64(len(products))
	for _, product := range products {
		repeat(
			func() error { return Buy(product, amount) },
			time.Second,
			100,
		)
	}

	return nil
}
