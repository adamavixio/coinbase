package coinbaseclient

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
