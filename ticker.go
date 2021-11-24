package coinbaseclient

import (
	"context"
	"encoding/json"

	logger "github.com/adamavixio/logger"
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

func StartTicker() *websocket.Conn {
	ctx := context.Background()
	client, _, err := websocket.Dial(ctx, socket, nil)
	logger.Error("coinbase websocket dial error", err)

	subscription, err := newSubscription().toJSON()
	logger.Error("subscription error", err)

	err = client.Write(ctx, websocket.MessageText, subscription)
	logger.Error("coinbase websocket write error", err)

	return client
}

func ReadTicker(client *websocket.Conn) *Ticker {
	_, bytes, err := client.Read(context.Background())
	logger.Error("coinbase websocket reading error", err)

	data := &Ticker{}
	err = json.Unmarshal(bytes, data)
	logger.Error("coinbase websocket reading error: %v", err)

	return data
}

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
