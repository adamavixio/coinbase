package client

import (
	"regexp"
	"strconv"
)

func MaxSize(productID string) (string, error) {
	singleProductConfig := &SingleProductConfig{
		ProductID: productID,
	}

	product, err := SingleProduct(singleProductConfig)
	if err != nil {
		return "", err
	}

	return product.BaseMaxSize, nil
}

func MinSize(productID string) (string, error) {
	singleProductConfig := &SingleProductConfig{
		ProductID: productID,
	}

	product, err := SingleProduct(singleProductConfig)
	if err != nil {
		return "", err
	}

	return product.BaseMinSize, nil
}

func Price(productID string) (float64, error) {
	config := &ProductTickerConfig{
		ProductID: productID,
	}

	ticker, err := ProductTicker(config)
	if err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(ticker.Price, 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}

func USDProductIDs() ([]string, error) {
	products, err := Products()
	if err != nil {
		return nil, err
	}

	regex := regexp.MustCompile(`USD$`)
	ids := []string{}

	for _, product := range products {
		if regex.MatchString(product.ID) {
			ids = append(ids, product.ID)
		}
	}

	return ids, nil
}
