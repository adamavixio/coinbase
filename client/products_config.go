package client

import (
	"fmt"
	"strconv"
)

type ProductTickerConfig struct {
	ProductID string `json:"product_id"`
}

func (config *ProductTickerConfig) validate() error {
	if config.ProductID == "" {
		return fmt.Errorf("invalid value for field ID: %s", config.ProductID)
	}

	return nil
}

type SingleProductConfig struct {
	ProductID string `json:"product_id"`
}

func (config *SingleProductConfig) validate() error {
	if config.ProductID == "" {
		return fmt.Errorf("invalid value for field ID: %s", config.ProductID)
	}

	return nil
}

type ProductTradesConfig struct {
	ProductID string `json:"product_id"`
	Limit     int    `json:"limit"`
	Before    int    `json:"before"`
	After     int    `json:"after"`
}

func (config *ProductTradesConfig) validate() error {
	if config.ProductID == "" {
		return fmt.Errorf("invalid value for field ID: %s", config.ProductID)
	}

	return nil
}

func (config *ProductTradesConfig) toMap() map[string]string {
	return map[string]string{
		"product_id": config.ProductID,
		"limit":      strconv.Itoa(config.Limit),
		"before":     strconv.Itoa(config.Before),
		"after":      strconv.Itoa(config.After),
	}
}
