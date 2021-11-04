package coinbaseapi

import (
	"context"
	"encoding/json"

	"nhooyr.io/websocket"
)

type Subscription struct {
	Type       string   `json:"type"`
	ProductIDs []string `json:"product_ids"`
	Channels   []string `json:"channels"`
}

func newSubscription() *Subscription {
	return &Subscription{
		Type:       "subscribe",
		ProductIDs: USDProductIDs(),
		Channels:   []string{"ticker"},
	}
}

func (s *Subscription) toJSON() ([]byte, error) {
	return json.Marshal(s)
}

type Ticker struct {
	Type      string `json:"type"`
	TradeId   int64  `json:"trade_id"`
	Sequence  int64  `json:"sequence"`
	Time      string `json:"time"`
	ProductId string `json:"product_id"`
	Price     string `json:"price"`
	Side      string `json:"side"`
	LastSize  string `json:"last_size"`
	BestBid   string `json:"best_bid"`
	BestAsk   string `json:"best_ask"`
}

func StartTicker() *websocket.Conn {
	ctx := context.Background()
	client, _, err := websocket.Dial(ctx, socket, nil)
	handleError("coinbase websocket dial error", err)

	subscription, err := newSubscription().toJSON()
	handleError("subscription error", err)

	if err := client.Write(ctx, websocket.MessageText, subscription); err != nil {
		handleError("coinbase websocket write error", err)
	}

	return client
}

func ReadTicker(client *websocket.Conn) *Ticker {
	_, bytes, err := client.Read(context.Background())
	handleError("coinbase websocket reading error", err)

	data := &Ticker{}
	if err := json.Unmarshal(bytes, data); err != nil {
		handleError("coinbase websocket reading error: %v", err)
	}

	return data
}
