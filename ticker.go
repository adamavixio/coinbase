package coinbaseclient

import (
	"context"
	"encoding/json"

	"nhooyr.io/websocket"
)

type Ticker struct {
	Type      string `json:"type" bson:"type"`
	TradeId   int64  `json:"trade_id" bson:"trade_id"`
	Sequence  int64  `json:"sequence" bson:"sequence"`
	Time      string `json:"time" bson:"time"`
	ProductId string `json:"product_id" bson:"product_id"`
	Price     string `json:"price" bson:"price"`
	Side      string `json:"side" bson:"side"`
	LastSize  string `json:"last_size" bson:"last_size"`
	BestBid   string `json:"best_bid" bson:"best_bid"`
	BestAsk   string `json:"best_ask" bson:"best_ask"`
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
