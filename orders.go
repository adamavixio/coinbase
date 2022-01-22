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

type GetOrders struct {
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

func (getOrders *GetOrders) isValid() error {
	if getOrders.Limit <= 0 {
		return errors.New("productID cannot be empty")
	}

	if _, ok := statuses[getOrders.Status]; !ok {
		return errors.New("status is not valid")
	}

	return nil
}

func (getOrders *GetOrders) toMap() map[string]string {
	return map[string]string{
		"after":      getOrders.After,
		"before":     getOrders.Before,
		"end_date":   getOrders.EndDate.String(),
		"limit":      fmt.Sprint(getOrders.Limit),
		"product_id": getOrders.ProductID,
		"profile_id": getOrders.ProfileID,
		"sortedBy":   getOrders.SortedBy,
		"sorting":    getOrders.Sorting,
		"start_date": getOrders.StartDate.String(),
		"status":     getOrders.Status,
	}
}

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

type GetOrder struct {
	OrderID string `json:"order_id" bson:"order_id"`
}

func (getOrder *GetOrder) isValid() error {
	if getOrder.OrderID == "" {
		return errors.New("orderID cannot be empty")
	}

	return nil
}

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

//
// API
//

func ExecuteGetOrders(getOrders *GetOrders) ([]*Order, error) {
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

func ExecuteCancelOrders(cancelOrders *CancelOrders) ([]string, error) {
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

func ExecuteCreateOrder(createOrder *CreateOrder) (*Order, error) {
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

func ExecuteGetOrder(getOrder *GetOrder) (*Order, error) {
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

func ExecuteCancelOrder(cancelOrder *CancelOrder) (*string, error) {
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
