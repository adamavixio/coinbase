package coinbaseclient

import (
	"encoding/json"
	"fmt"
	"regexp"
)

//
// Implementation
//

type Product struct {
	AuctionMode           bool   `json:"auction_mode,omitempty" bson:"auction_mode,omitempty"`
	BaseCurrency          string `json:"base_currency,omitempty" bson:"base_currency,omitempty"`
	BaseIncrement         string `json:"base_increment,omitempty" bson:"base_increment,omitempty"`
	BaseMaxSize           string `json:"base_max_size,omitempty" bson:"base_max_size,omitempty"`
	BaseMinSize           string `json:"base_min_size,omitempty" bson:"base_min_size,omitempty"`
	CancelOnly            bool   `json:"cancel_only,omitempty" bson:"cancel_only,omitempty"`
	DisplayName           string `json:"display_name,omitempty" bson:"display_name,omitempty"`
	FXStableCoin          bool   `json:"fx_stablecoin,omitempty" bson:"fx_stablecoin,omitempty"`
	ID                    string `json:"id,omitempty" bson:"id,omitempty"`
	LimitOnly             bool   `json:"limit_only,omitempty" bson:"limit_only,omitempty"`
	MarginEnabled         bool   `json:"margin_enabled,omitempty" bson:"margin_enabled,omitempty"`
	MaxSlippagePercentage string `json:"max_slippage_percentage,omitempty" bson:"max_slippage_percentage,omitempty"`
	MaxMarketFunds        string `json:"max_market_funds,omitempty" bson:"max_market_funds,omitempty"`
	MinMarketFunds        string `json:"min_market_funds,omitempty" bson:"min_market_funds,omitempty"`
	PostOnly              bool   `json:"post_only,omitempty" bson:"post_only,omitempty"`
	QuoteCurrency         string `json:"quote_currency,omitempty" bson:"quote_currency,omitempty"`
	QuoteIncrement        string `json:"quote_increment,omitempty" bson:"quote_increment,omitempty"`
	Status                string `json:"status,omitempty" bson:"status,omitempty"`
	StatusMessage         string `json:"status_message,omitempty" bson:"status_message,omitempty"`
	TradingDisabled       bool   `json:"trading_disabled,omitempty" bson:"trading_disabled,omitempty"`
}

type ProductTickerConfig struct {
	ID string `json:"product_id"`
}

func (config *ProductTickerConfig) isValid() error {
	if config.ID == "" {
		return fmt.Errorf("invalid value for field ID: %s", config.ID)
	}

	return nil
}

//
// API
//

func Products() ([]*Product, error) {
	config := RequestConfig{
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

func SingleProduct(config *ProductTickerConfig) (*Ticker, error) {
	err := config.isValid()
	if err != nil {
		return nil, err
	}

	request := RequestConfig{
		Method: get,
		Path:   fmt.Sprintf("/products/%s/ticker", config.ID),
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
