package coinbaseclient

import (
	"encoding/json"
	"errors"
	"fmt"
)

//
// Defintions
//

type CancelOrder struct {
	OrderID   string `json:"order_id" bson:"order_id"`
	ProfileID string `json:"profile_id,omitempty" bson:"profile_id,omitempty"`
}

func (cancelOrder *CancelOrder) isValid() error {
	if cancelOrder.OrderID == "" {
		return errors.New("OrderID cannot be empty")
	}

	return nil
}

func (cancelOrder *CancelOrder) toMap() map[string]string {
	return map[string]string{
		"profile_id": cancelOrder.ProfileID,
	}
}

type CanceledOrder string

//
// Exports
//

func ExecuteCancelOrder(cancelOrder *CancelOrder) (*CanceledOrder, error) {
	err := cancelOrder.isValid()
	if err != nil {
		return nil, err
	}

	config := RequestConfig{
		Method: delete,
		Path:   fmt.Sprintf("/orders/%s", cancelOrder.OrderID),
		Params: cancelOrder.toMap(),
	}

	data, err := executeAuthenticatedRequest(config)
	if err != nil {
		return nil, err
	}

	var canceledOrder *CanceledOrder

	err = json.Unmarshal(data, canceledOrder)
	if err != nil {
		return nil, err
	}

	return canceledOrder, err
}
