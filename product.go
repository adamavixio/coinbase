package coinbaseclient

import (
	"encoding/json"
	"regexp"
)

type Product struct {
	ID                    string `json:"id,omitempty" bson:"id,omitempty"`
	BaseCurrency          string `json:"base_currency,omitempty" bson:"base_currency,omitempty"`
	QuoteCurrency         string `json:"quote_currency,omitempty" bson:"quote_currency,omitempty"`
	BaseMinSize           string `json:"base_min_size,omitempty" bson:"base_min_size,omitempty"`
	BaseMaxSize           string `json:"base_max_size,omitempty" bson:"base_max_size,omitempty"`
	QuoteIncrement        string `json:"quote_increment,omitempty" bson:"quote_increment,omitempty"`
	BaseIncrement         string `json:"base_increment,omitempty" bson:"base_increment,omitempty"`
	DisplayName           string `json:"display_name,omitempty" bson:"display_name,omitempty"`
	MinMarketFunds        string `json:"min_market_funds,omitempty" bson:"min_market_funds,omitempty"`
	MaxMarketFunds        string `json:"max_market_funds,omitempty" bson:"max_market_funds,omitempty"`
	MarginEnabled         bool   `json:"margin_enabled,omitempty" bson:"margin_enabled,omitempty"`
	FXStableCoin          bool   `json:"fx_stablecoin,omitempty" bson:"fx_stablecoin,omitempty"`
	MaxSlippagePercentage string `json:"max_slippage_percentage,omitempty" bson:"max_slippage_percentage,omitempty"`
	PostOnly              bool   `json:"post_only,omitempty" bson:"post_only,omitempty"`
	LimitOnly             bool   `json:"limit_only,omitempty" bson:"limit_only,omitempty"`
	CancelOnly            bool   `json:"cancel_only,omitempty" bson:"cancel_only,omitempty"`
	TradingDisabled       bool   `json:"trading_disabled,omitempty" bson:"trading_disabled,omitempty"`
	Status                string `json:"status,omitempty" bson:"status,omitempty"`
	StatusMessage         string `json:"status_message,omitempty" bson:"status_message,omitempty"`
	AuctionMode           bool   `json:"auction_mode,omitempty" bson:"auction_mode,omitempty"`
}

func Products() ([]Product, error) {
	config := RequestConfig{
		Method: get,
		Path:   "/products",
	}

	data, err := executeAuthenticatedRequest(config)
	if err != nil {
		return nil, err
	}

	products := []Product{}

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
