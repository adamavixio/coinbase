package client

import (
	"encoding/json"
)

type Subscription struct {
	Channels   []string `json:"channels"`
	ProductIDs []string `json:"product_ids"`
	Type       string   `json:"type"`
}

func newSubscription() (*Subscription, error) {
	productIDs, err := USDProductIDs()
	if err != nil {
		return nil, err
	}

	subscription := &Subscription{
		Type:       "subscribe",
		ProductIDs: productIDs,
		Channels:   []string{"ticker"},
	}

	return subscription, nil
}

func (s *Subscription) toJSON() ([]byte, error) {
	return json.Marshal(s)
}
