package service

import "testing"

// func Test_Buy(t *testing.T) {
// 	productID, amount := "BTC-USD", float64(100)
// 	err := Buy(productID, amount)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func Test_BuyAll(t *testing.T) {
// 	productID := "BTC-USD"
// 	err := BuyAll(productID)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func Test_Sell(t *testing.T) {
// 	productID, amount := "BTC-USD", float64(0.003)
// 	err := Sell(productID, amount)
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

func Test_SellAll(t *testing.T) {
	productID := "BTC-USD"
	err := SellAll(productID)
	if err != nil {
		t.Error(err)
	}
}

// func Test_Spread(t *testing.T) {
// 	err := Spread()
// 	if err != nil {
// 		t.Error(err)
// 	}
// }
