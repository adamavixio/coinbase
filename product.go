package coinbaseapi

import (
	"encoding/json"
	"strings"
)

type Product struct {
	ID                    string `json:"id"`
	BaseCurrency          string `json:"base_currency"`
	QuoteCurrency         string `json:"quote_currency"`
	BaseMinSize           string `json:"base_min_size"`
	BaseMaxSize           string `json:"base_max_size"`
	QuoteIncrement        string `json:"quote_increment"`
	BaseIncrement         string `json:"base_increment"`
	DisplayName           string `json:"display_name"`
	MinMarketFunds        string `json:"min_market_funds"`
	MaxMarketFunds        string `json:"max_market_funds"`
	MarginEnabled         bool   `json:"margin_enabled"`
	FXStableCoin          bool   `json:"fx_stablecoin"`
	MaxSlippagePercentage string `json:"max_slippage_percentage"`
	PostOnly              bool   `json:"post_only"`
	LimitOnly             bool   `json:"limit_only"`
	CancelOnly            bool   `json:"cancel_only"`
	TradingDisabled       bool   `json:"trading_disabled"`
	Status                string `json:"status"`
	StatusMessage         string `json:"status_message"`
	AuctionMode           bool   `json:"auction_mode"`
}

func Products() []Product {
	data := executeAuthenticatedRequest(get, "/products", nil, nil)

	products := []Product{}
	if err := json.Unmarshal(data, &products); err != nil {
		handleError("coinbase product unmarshal error", err)
	}

	return products
}

func USDProductIDs() []string {
	products := Products()

	ids := []string{}
	for _, product := range products {
		if strings.Contains(product.ID, "USD") {
			ids = append(ids, product.ID)
		}
	}

	return ids
}
