package coinbaseclient

import (
	"context"
	"encoding/json"

	"nhooyr.io/websocket"
)

type Ticker struct {
	Type      string `json:"type,omitempty" bson:"type,omitempty"`
	TradeId   int64  `json:"trade_id,omitempty" bson:"trade_id,omitempty"`
	Sequence  int64  `json:"sequence,omitempty" bson:"sequence,omitempty"`
	Time      string `json:"time,omitempty" bson:"time,omitempty"`
	ProductId string `json:"product_id,omitempty" bson:"product_id,omitempty"`
	Price     string `json:"price,omitempty" bson:"price,omitempty"`
	Side      string `json:"side,omitempty" bson:"side,omitempty"`
	LastSize  string `json:"last_size,omitempty" bson:"last_size,omitempty"`
	BestBid   string `json:"best_bid,omitempty" bson:"best_bid,omitempty"`
	BestAsk   string `json:"best_ask,omitempty" bson:"best_ask,omitempty"`
}

func StartTicker() (*websocket.Conn, error) {
	ctx := context.Background()

	client, _, err := websocket.Dial(ctx, socket, nil)
	if err != nil {
		return nil, err
	}

	subscription, err := newSubscription()
	if err != nil {
		return nil, err
	}

	bytes, err := subscription.toJSON()
	if err != nil {
		return nil, err
	}

	err = client.Write(ctx, websocket.MessageText, bytes)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func ReadTicker(client *websocket.Conn) (*Ticker, error) {
	_, bytes, err := client.Read(context.Background())
	if err != nil {
		return nil, err
	}

	data := &Ticker{}

	err = json.Unmarshal(bytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type Subscription struct {
	Type       string   `json:"type"`
	ProductIDs []string `json:"product_ids"`
	Channels   []string `json:"channels"`
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
