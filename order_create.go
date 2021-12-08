package coinbaseclient

import (
	"encoding/json"
	"errors"
)

//
// Definitions
//

type CreateOrder struct {
	CancelAfter string `json:"cancel_after,omitempty" bson:"cancel_after,omitempty"`
	ClientOID   string `json:"client_oid,omitempty" bson:"client_oid,omitempty"`
	Funds       string `json:"funds,omitempty" bson:"funds,omitempty"`
	PostOnly    bool   `json:"post_only,omitempty" bson:"post_only,omitempty"`
	Price       string `json:"price,omitempty" bson:"price,omitempty"`
	ProductID   string `json:"product_id" bson:"product_id"`
	ProfileID   string `json:"profile_id,omitempty" bson:"profile_id,omitempty"`
	Side        string `json:"side" bson:"side"`
	Size        string `json:"size,omitempty" bson:"size,omitempty"`
	Stop        string `json:"stop,omitempty" bson:"stop,omitempty"`
	StopPrice   string `json:"stop_price,omitempty" bson:"stop_price,omitempty"`
	Stp         string `json:"stp,omitempty" bson:"stp,omitempty"`
	TimeInForce string `json:"time_in_force,omitempty" bson:"time_in_force,omitempty"`
	Type        string `json:"type,omitempty" bson:"type,omitempty"`
}

func (createOrder *CreateOrder) isValid() error {
	if createOrder.ProductID == "" {
		return errors.New("ProductID cannot be empty")
	}

	return nil
}

func (createOrder *CreateOrder) toJSON() ([]byte, error) {
	return json.Marshal(createOrder)
}

type CreatedOrder struct {
	CreatedAt     string `json:"created_at,omitempty" bson:"created_at,omitempty"`
	ExecutedValue string `json:"executed_value,omitempty" bson:"executed_value,omitempty"`
	FillFees      string `json:"fill_fees,omitempty" bson:"fill_fees,omitempty"`
	FilledSize    string `json:"filled_size,omitempty" bson:"filled_size,omitempty"`
	ID            string `json:"id,omitempty" bson:"id,omitempty"`
	PostOnly      bool   `json:"post_only,omitempty" bson:"post_only,omitempty"`
	Price         string `json:"price,omitempty" bson:"price,omitempty"`
	ProductID     string `json:"product_id,omitempty" bson:"product_id,omitempty"`
	ProfileID     string `json:"profile_id,omitempty" bson:"profile_id,omitempty"`
	Settled       bool   `json:"settled,omitempty" bson:"settled,omitempty"`
	Side          string `json:"side,omitempty" bson:"side,omitempty"`
	Size          string `json:"size,omitempty" bson:"size,omitempty"`
	Status        string `json:"status,omitempty" bson:"status,omitempty"`
	TimeInForce   string `json:"time_in_force,omitempty" bson:"time_in_force,omitempty"`
	Type          string `json:"type,omitempty" bson:"type,omitempty"`
}

//
// Exports
//

func ExecuteCreateOrder(createOrder *CreateOrder) (*CreatedOrder, error) {
	err := createOrder.isValid()
	if err != nil {
		return nil, err
	}

	bytes, err := createOrder.toJSON()
	if err != nil {
		return nil, err
	}

	config := RequestConfig{
		Method: post,
		Path:   "/orders",
		Body:   bytes,
	}

	data, err := executeAuthenticatedRequest(config)
	if err != nil {
		return nil, err
	}

	createdOrder := CreatedOrder{}

	err = json.Unmarshal(data, &createdOrder)
	if err != nil {
		return nil, err
	}

	return &createdOrder, err
}
