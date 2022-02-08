package client

import (
	"encoding/json"
	"fmt"
)

var (
	statuses = map[string]struct{}{
		"active":   {},
		"all":      {},
		"done":     {},
		"open":     {},
		"pending":  {},
		"received": {},
		"rejected": {},
	}
)

type Order struct {
	CreatedAt     string `json:"created_at" bson:"created_at"`
	ExecutedValue string `json:"executed_value" bson:"executed_value"`
	FillFees      string `json:"fill_fees" bson:"fill_fees"`
	FilledSize    string `json:"filled_size" bson:"filled_size"`
	OrderID       string `json:"order_id" bson:"order_id"`
	PostOnly      bool   `json:"post_only" bson:"post_only"`
	Price         string `json:"price" bson:"price"`
	ProductID     string `json:"product_id" bson:"product_id"`
	ProfileID     string `json:"profile_id" bson:"profile_id"`
	Settled       bool   `json:"settled" bson:"settled"`
	Side          string `json:"side" bson:"side"`
	Size          string `json:"size" bson:"size"`
	Status        string `json:"status" bson:"status"`
	TimeInForce   string `json:"time_in_force" bson:"time_in_force"`
	Type          string `json:"type" bson:"type"`
}

func GetOrders(config *GetOrdersConfig) ([]*Order, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	request := requestConfig{
		Method: get,
		Path:   "/orders",
		Params: config.toMap(),
	}

	data, err := authRequest(request)
	if err != nil {
		return nil, err
	}

	orders := []*Order{}

	err = json.Unmarshal(data, &orders)
	if err != nil {
		return nil, err
	}

	return orders, err
}

func CancelOrders(config *CancelOrdersConfig) ([]string, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	request := requestConfig{
		Method: delete,
		Path:   "/orders",
		Params: config.toMap(),
	}

	data, err := authRequest(request)
	if err != nil {
		return nil, err
	}

	canceledOrders := []string{}

	err = json.Unmarshal(data, &canceledOrders)
	if err != nil {
		return nil, err
	}

	return canceledOrders, err
}

func CreateOrder(config *CreateOrderConfig) (*Order, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	bytes, err := config.toJSON()
	if err != nil {
		return nil, err
	}

	request := requestConfig{
		Method: post,
		Path:   "/orders",
		Body:   bytes,
	}

	data, err := authRequest(request)
	if err != nil {
		return nil, err
	}

	createdOrder := Order{}

	err = json.Unmarshal(data, &createdOrder)
	if err != nil {
		return nil, err
	}

	return &createdOrder, err
}

func GetOrder(config *GetOrderConfig) (*Order, error) {
	if err := config.validate(); err != nil {
		return nil, err
	}

	request := requestConfig{
		Method: get,
		Path:   fmt.Sprintf("/orders/%s", config.OrderID),
	}

	data, err := authRequest(request)
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

func CancelOrder(cancelOrder *CancelOrderConfig) (*string, error) {
	if err := cancelOrder.validate(); err != nil {
		return nil, err
	}

	config := requestConfig{
		Method: delete,
		Path:   fmt.Sprintf("/orders/%s", cancelOrder.OrderID),
		Params: cancelOrder.toMap(),
	}

	data, err := authRequest(config)
	if err != nil {
		return nil, err
	}

	var canceledOrder *string

	err = json.Unmarshal(data, canceledOrder)
	if err != nil {
		return nil, err
	}

	return canceledOrder, err
}
