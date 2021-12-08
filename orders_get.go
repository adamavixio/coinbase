package coinbaseclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

//
// Definitions
//

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

//
// Exports
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

	data, err := executeAuthenticatedRequest(config)
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
