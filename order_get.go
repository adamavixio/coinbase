package coinbaseclient

import (
	"encoding/json"
	"errors"
	"fmt"
)

//
// Definitions
//

type GetOrder struct {
	OrderID string `json:"order_id" bson:"order_id"`
}

func (getOrder *GetOrder) isValid() error {
	if getOrder.OrderID == "" {
		return errors.New("orderID cannot be empty")
	}

	return nil
}

//
// Exports
//

func ExecuteGetOrder(getOrder *GetOrder) (*Order, error) {
	err := getOrder.isValid()
	if err != nil {
		return nil, err
	}

	config := RequestConfig{
		Method: get,
		Path:   fmt.Sprintf("/orders/%s", getOrder.OrderID),
	}

	data, err := executeAuthenticatedRequest(config)
	if err != nil {
		return nil, err
	}

	order := &Order{}

	err = json.Unmarshal(data, order)
	if err != nil {
		return nil, err
	}

	return order, err
}
