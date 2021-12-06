package coinbaseclient

import (
	"encoding/json"
	"fmt"

	logger "github.com/adamavixio/logger"
)

type Trade struct {
	Time    string `json:"time"`
	TradeID int64  `json:"trade_id"`
	Price   string `json:"price"`
	Size    string `json:"size"`
	Side    string `json:"side"`
}

func Trades(productID string, before string) []Trade {
	params := map[string]string{}
	params["before"] = before

	path := fmt.Sprintf("/products/%s/trades", productID)
	data := executeAuthenticatedRequest(get, path, params, nil)

	trades := []Trade{}
	err := json.Unmarshal(data, &trades)
	logger.Error("coinbase product unmarshal error", err)

	return trades
}
