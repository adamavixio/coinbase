package client

func Orders() ([]*Order, error) {
	config := &GetOrdersConfig{
		Limit:  10000,
		Status: "done",
	}

	orders, err := GetOrders(config)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func Buy(productID string, size string, orderType string) (*Order, error) {
	config := &CreateOrderConfig{
		ProductID: productID,
		Size:      size,
		Side:      "buy",
		Type:      orderType,
	}

	order, err := CreateOrder(config)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func Sell(productID string, size string, orderType string) (*Order, error) {
	config := &CreateOrderConfig{
		ProductID: productID,
		Size:      size,
		Side:      "sell",
		Type:      orderType,
	}

	order, err := CreateOrder(config)
	if err != nil {
		return nil, err
	}

	return order, nil
}
