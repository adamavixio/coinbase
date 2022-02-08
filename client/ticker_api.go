package client

import (
	"context"
	"encoding/json"

	"nhooyr.io/websocket"
)

type Ticker struct {
	BestAsk   string `json:"best_ask" bson:"best_ask"`
	BestBid   string `json:"best_bid" bson:"best_bid"`
	LastSize  string `json:"last_size" bson:"last_size"`
	Price     string `json:"price" bson:"price"`
	ProductId string `json:"product_id" bson:"product_id"`
	Sequence  int64  `json:"sequence" bson:"sequence"`
	Side      string `json:"side" bson:"side"`
	Time      string `json:"time" bson:"time"`
	TradeId   int64  `json:"trade_id" bson:"trade_id"`
	Type      string `json:"type" bson:"type"`
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
