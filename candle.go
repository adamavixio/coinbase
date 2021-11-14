package coinbaseapi

import (
	"encoding/json"
	"fmt"
)

type Candles = []TLHOCV
type TLHOCV = []int

func CandleSticks(product, granularity, start, end string) Candles {
	params := map[string]string{
		"product":     product,
		"granularity": granularity,
		"start":       start,
		"end":         end,
	}

	path := fmt.Sprintf("/products/%s/candles", product)
	data := executeAuthenticatedRequest(get, path, params, nil)

	candles := Candles{}
	if err := json.Unmarshal(data, &candles); err != nil {
		handleError("coinbase product unmarshal error", err)
	}

	return candles
}
