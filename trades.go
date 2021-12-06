package coinbaseclient

import (
	"encoding/json"
	"fmt"
)

type Trade struct {
	Time    string `json:"time" bson:"time"`
	TradeID int64  `json:"trade_id" bson:"trade_id"`
	Price   string `json:"price" bson:"price"`
	Size    string `json:"size" bson:"size"`
	Side    string `json:"side" bson:"side"`
}

func Trades(productID string, after string) ([]Trade, error) {
	config := RequestConfig{
		Method: get,
		Path:   fmt.Sprintf("/products/%s/trades", productID),
		Params: map[string]string{"after": after},
	}

	data, err := executeAuthenticatedRequest(config)
	if err != nil {
		return nil, err
	}

	trades := []Trade{}

	err = json.Unmarshal(data, &trades)
	if err != nil {
		return nil, err
	}

	return trades, err
}
