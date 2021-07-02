package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
)

type CurrencyQuote struct {
	Price decimal.Decimal `json:"price"`
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

func (d *CoinmarketcapQuotesRateDecoder) Decode(raw []byte, from, to string) (decimal.Decimal, error) {

	var r Response
	err := json.Unmarshal(raw, &r)

	if err != nil {
		return decimal.Zero, fmt.Errorf("Unmarshalling error: %v", err)
	}

	if r.Status.ErrorCode != 0 {
		return decimal.Zero, fmt.Errorf("Bad Status, error_code:%v error_message:%v", r.Status.ErrorCode, r.Status.ErrorMessage)
	}

	entry, ok := r.Entries[to]
	if !ok {
		return decimal.Zero, fmt.Errorf("No entry for currency: %v", to)
	}

	quote, ok := entry.CurrencyQuotes[from]
	if !ok {
		return decimal.Zero, fmt.Errorf("No quote for currency: %v", from)
	}

	return quote.Price, nil
}
