package coinbaseclient

import (
	"encoding/json"
	"errors"
)

//
// Definitions
//

type CancelOrders struct {
	ProductID string `json:"product_id,omitempty" bson:"product_id,omitempty"`
	ProfileID string `json:"profile_id,omitempty" bson:"profile_id,omitempty"`
}

func (cancelOrders *CancelOrders) isValid() error {
	if cancelOrders.ProductID == "" {
		return errors.New("ProductID cannot be empty")
	}

	return nil
}

func (cancelOrders *CancelOrders) toMap() map[string]string {
	return map[string]string{
		"profile_id": cancelOrders.ProfileID,
		"product_id": cancelOrders.ProductID,
	}
}

type CanceledOrders []string

//
// Exports
//

func ExecuteCancelOrders(cancelOrders *CancelOrders) (CanceledOrders, error) {
	err := cancelOrders.isValid()
	if err != nil {
		return nil, err
	}

	config := RequestConfig{
		Method: delete,
		Path:   "/orders",
		Params: cancelOrders.toMap(),
	}

	data, err := executeAuthenticatedRequest(config)
	if err != nil {
		return nil, err
	}

	canceledOrders := CanceledOrders{}

	err = json.Unmarshal(data, &canceledOrders)
	if err != nil {
		return nil, err
	}

	return canceledOrders, err
}
