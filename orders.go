package coinbaseclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
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

//
// Implementation
//

type Order struct {
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

type GetOrdersConfig struct {
	After     string    `json:"after,omitempty" bson:"after,omitempty"`
	Before    string    `json:"before,omitempty" bson:"before,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty" bson:"end_date,omitempty"`
	Limit     int64     `json:"limit" bson:"limit"`
	ProductID string    `json:"product_id,omitempty" bson:"product_id,omitempty"`
	ProfileID string    `json:"profile_id,omitempty" bson:"profile_id,omitempty"`
	SortedBy  string    `json:"sortedBy,omitempty" bson:"sortedBy,omitempty"`
	Sorting   string    `json:"sorting,omitempty" bson:"sorting,omitempty"`
	StartDate time.Time `json:"start_date,omitempty" bson:"start_date,omitempty"`
	Status    string    `json:"status" bson:"status"`
}

func (config *GetOrdersConfig) isValid() error {
	if config.Limit <= 0 {
		return errors.New("productID cannot be empty")
	}

	if _, ok := statuses[config.Status]; !ok {
		return errors.New("status is not valid")
	}

	return nil
}

func (config *GetOrdersConfig) toMap() map[string]string {
	return map[string]string{
		"after":      config.After,
		"before":     config.Before,
		"end_date":   config.EndDate.String(),
		"limit":      fmt.Sprint(config.Limit),
		"product_id": config.ProductID,
		"profile_id": config.ProfileID,
		"sortedBy":   config.SortedBy,
		"sorting":    config.Sorting,
		"start_date": config.StartDate.String(),
		"status":     config.Status,
	}
}

type CancelOrdersConfig struct {
	ProductID string `json:"product_id,omitempty" bson:"product_id,omitempty"`
	ProfileID string `json:"profile_id,omitempty" bson:"profile_id,omitempty"`
}

func (config *CancelOrdersConfig) isValid() error {
	if config.ProductID == "" {
		return errors.New("ProductID cannot be empty")
	}

	return nil
}

func (config *CancelOrdersConfig) toMap() map[string]string {
	return map[string]string{
		"profile_id": config.ProfileID,
		"product_id": config.ProductID,
	}
}

type CreateOrderConfig struct {
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

func (config *CreateOrderConfig) isValid() error {
	if config.ProductID == "" {
		return errors.New("ProductID cannot be empty")
	}

	return nil
}

func (config *CreateOrderConfig) toJSON() ([]byte, error) {
	return json.Marshal(config)
}

type GetOrderConfig struct {
	OrderID string `json:"order_id" bson:"order_id"`
}

func (config *GetOrderConfig) isValid() error {
	if config.OrderID == "" {
		return errors.New("orderID cannot be empty")
	}

	return nil
}

type CancelOrderConfig struct {
	OrderID   string `json:"order_id" bson:"order_id"`
	ProfileID string `json:"profile_id,omitempty" bson:"profile_id,omitempty"`
}

func (config *CancelOrderConfig) isValid() error {
	if config.OrderID == "" {
		return errors.New("OrderID cannot be empty")
	}

	return nil
}

func (config *CancelOrderConfig) toMap() map[string]string {
	return map[string]string{
		"profile_id": config.ProfileID,
	}
}

//
// API
//

func ExecuteGetOrders(getOrders *GetOrdersConfig) ([]*Order, error) {
	err := getOrders.isValid()
	if err != nil {
		return nil, err
	}

	config := RequestConfig{
		Method: get,
		Path:   "/orders",
		Params: getOrders.toMap(),
	}

	data, err := authRequest(config)
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

func ExecuteCancelOrders(cancelOrders *CancelOrdersConfig) ([]string, error) {
	err := cancelOrders.isValid()
	if err != nil {
		return nil, err
	}

	config := RequestConfig{
		Method: delete,
		Path:   "/orders",
		Params: cancelOrders.toMap(),
	}

	data, err := authRequest(config)
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

func ExecuteCreateOrder(createOrder *CreateOrderConfig) (*Order, error) {
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

	data, err := authRequest(config)
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

func ExecuteGetOrder(getOrder *GetOrderConfig) (*Order, error) {
	err := getOrder.isValid()
	if err != nil {
		return nil, err
	}

	config := RequestConfig{
		Method: get,
		Path:   fmt.Sprintf("/orders/%s", getOrder.OrderID),
	}

	data, err := authRequest(config)
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

func ExecuteCancelOrder(cancelOrder *CancelOrderConfig) (*string, error) {
	err := cancelOrder.isValid()
	if err != nil {
		return nil, err
	}

	config := RequestConfig{
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
