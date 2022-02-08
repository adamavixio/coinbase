package client

import (
	"encoding/json"
	"fmt"
	"time"
)

type GetOrdersConfig struct {
	After     string    `json:"after" bson:"after"`
	Before    string    `json:"before" bson:"before"`
	EndDate   time.Time `json:"end_date" bson:"end_date"`
	Limit     int64     `json:"limit" bson:"limit"`
	ProductID string    `json:"product_id" bson:"product_id"`
	ProfileID string    `json:"profile_id" bson:"profile_id"`
	SortedBy  string    `json:"sortedBy" bson:"sortedBy"`
	Sorting   string    `json:"sorting" bson:"sorting"`
	StartDate time.Time `json:"start_date" bson:"start_date"`
	Status    string    `json:"status" bson:"status"`
}

func (config *GetOrdersConfig) validate() error {
	if config.Limit <= 0 {
		return fmt.Errorf("productID cannot be empty")
	}

	if _, ok := statuses[config.Status]; !ok {
		return fmt.Errorf("status is not valid")
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
	ProductID string `json:"product_id" bson:"product_id"`
	ProfileID string `json:"profile_id" bson:"profile_id"`
}

func (config *CancelOrdersConfig) validate() error {
	if config.ProductID == "" {
		return fmt.Errorf("ProductID cannot be empty")
	}

	return nil
}

func (config *CancelOrdersConfig) toMap() map[string]string {
	return map[string]string{
		"product_id": config.ProductID,
		"profile_id": config.ProfileID,
	}
}

type CreateOrderConfig struct {
	CancelAfter string `json:"cancel_after" bson:"cancel_after"`
	ClientOID   string `json:"client_oid" bson:"client_oid"`
	Funds       string `json:"funds" bson:"funds"`
	PostOnly    bool   `json:"post_only" bson:"post_only"`
	Price       string `json:"price" bson:"price"`
	ProductID   string `json:"product_id" bson:"product_id"`
	ProfileID   string `json:"profile_id" bson:"profile_id"`
	Side        string `json:"side" bson:"side"`
	Size        string `json:"size" bson:"size"`
	Stop        string `json:"stop" bson:"stop"`
	StopPrice   string `json:"stop_price" bson:"stop_price"`
	Stp         string `json:"stp" bson:"stp"`
	TimeInForce string `json:"time_in_force" bson:"time_in_force"`
	Type        string `json:"type" bson:"type"`
}

func (config *CreateOrderConfig) validate() error {
	if config.ProductID == "" {
		return fmt.Errorf("productID cannot be empty")
	}

	if config.Side == "" {
		return fmt.Errorf("side cannot be empty")
	}

	return nil
}

func (config *CreateOrderConfig) toJSON() ([]byte, error) {
	return json.Marshal(config)
}

type GetOrderConfig struct {
	OrderID string `json:"order_id" bson:"order_id"`
}

func (config *GetOrderConfig) validate() error {
	if config.OrderID == "" {
		return fmt.Errorf("orderID cannot be empty")
	}

	return nil
}

type CancelOrderConfig struct {
	OrderID   string `json:"order_id" bson:"order_id"`
	ProfileID string `json:"profile_id" bson:"profile_id"`
}

func (config *CancelOrderConfig) validate() error {
	if config.OrderID == "" {
		return fmt.Errorf("OrderID cannot be empty")
	}

	return nil
}

func (config *CancelOrderConfig) toMap() map[string]string {
	return map[string]string{
		"profile_id": config.ProfileID,
	}
}
