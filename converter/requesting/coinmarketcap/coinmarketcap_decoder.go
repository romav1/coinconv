package coinmarketcap

import (
	"encoding/json"
	"fmt"
)

type CurrencyQuote struct {
	Price float64 `json:"price"`
}

type Entry struct {
	CurrencyQuotes map[string]CurrencyQuote `json:"quote"`
}

type Status struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type Response struct {
	Status  Status           `json:"status"`
	Entries map[string]Entry `json:"data"`
}

type CoinmarketcapQuotesRateDecoder struct {
}

func (d *CoinmarketcapQuotesRateDecoder) Decode(raw []byte, from, to string) (float64, error) {

	var r Response
	err := json.Unmarshal(raw, &r)

	if err != nil {
		return 0, fmt.Errorf("Unmarshalling error: %v", err)
	}

	if r.Status.ErrorCode != 0 {
		return 0, fmt.Errorf("Bad Status, error_code:%v error_message:%v", r.Status.ErrorCode, r.Status.ErrorMessage)
	}

	entry, ok := r.Entries[to]
	if !ok {
		return 0, fmt.Errorf("No entry for currency: %v", to)
	}

	quote, ok := entry.CurrencyQuotes[from]
	if !ok {
		return 0, fmt.Errorf("No quote for currency: %v", from)
	}

	return quote.Price, nil
}
