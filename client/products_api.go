package client

import (
	"encoding/json"
	"fmt"
)

type Trade struct {
	Time    string `json:"time,omitempty" bson:"time,omitempty"`
	TradeID int64  `json:"trade_id,omitempty" bson:"trade_id,omitempty"`
	Price   string `json:"price,omitempty" bson:"price,omitempty"`
	Size    string `json:"size,omitempty" bson:"size,omitempty"`
	Side    string `json:"side,omitempty" bson:"side,omitempty"`
}

type Product struct {
	AuctionMode           bool   `json:"auction_mode" bson:"auction_mode"`
	BaseCurrency          string `json:"base_currency" bson:"base_currency"`
	BaseIncrement         string `json:"base_increment" bson:"base_increment"`
	BaseMaxSize           string `json:"base_max_size" bson:"base_max_size"`
	BaseMinSize           string `json:"base_min_size" bson:"base_min_size"`
	CancelOnly            bool   `json:"cancel_only" bson:"cancel_only"`
	DisplayName           string `json:"display_name" bson:"display_name"`
	FXStableCoin          bool   `json:"fx_stablecoin" bson:"fx_stablecoin"`
	ID                    string `json:"id" bson:"id"`
	LimitOnly             bool   `json:"limit_only" bson:"limit_only"`
	MarginEnabled         bool   `json:"margin_enabled" bson:"margin_enabled"`
	MaxSlippagePercentage string `json:"max_slippage_percentage" bson:"max_slippage_percentage"`
	MaxMarketFunds        string `json:"max_market_funds" bson:"max_market_funds"`
	MinMarketFunds        string `json:"min_market_funds" bson:"min_market_funds"`
	PostOnly              bool   `json:"post_only" bson:"post_only"`
	QuoteCurrency         string `json:"quote_currency" bson:"quote_currency"`
	QuoteIncrement        string `json:"quote_increment" bson:"quote_increment"`
	Status                string `json:"status" bson:"status"`
	StatusMessage         string `json:"status_message" bson:"status_message"`
	TradingDisabled       bool   `json:"trading_disabled" bson:"trading_disabled"`
}

func Products() ([]*Product, error) {
	config := requestConfig{
		Method: get,
		Path:   "/products",
	}

	data, err := authRequest(config)
	if err != nil {
		return nil, err
	}

	products := []*Product{}

	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil, err
	}

	return products, err
}

func SingleProduct(config *SingleProductConfig) (*Product, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	request := requestConfig{
		Method: get,
		Path:   fmt.Sprintf("/products/%s", config.ProductID),
	}

	data, err := authRequest(request)
	if err != nil {
		return nil, err
	}

	product := &Product{}
	err = json.Unmarshal(data, product)
	if err != nil {
		return nil, err
	}

	return product, err
}

func ProductTicker(config *ProductTickerConfig) (*Ticker, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	request := requestConfig{
		Method: get,
		Path:   fmt.Sprintf("/products/%s/ticker", config.ProductID),
	}

	data, err := authRequest(request)
	if err != nil {
		return nil, err
	}

	ticker := &Ticker{}
	err = json.Unmarshal(data, ticker)
	if err != nil {
		return nil, err
	}

	return ticker, err
}

func ProductTrades(config *ProductTradesConfig) ([]*Trade, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	request := requestConfig{
		Method: get,
		Path:   fmt.Sprintf("/products/%s/trades", config.ProductID),
		Params: config.toMap(),
	}

	data, err := authRequest(request)
	if err != nil {
		return nil, err
	}

	trades := []*Trade{}

	err = json.Unmarshal(data, &trades)
	if err != nil {
		return nil, err
	}

	return trades, err
}
