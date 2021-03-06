package requesting

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type RateRequester interface {
	Request(from, to string) ([]byte, error)
}

type RateDecoder interface {
	Decode(raw []byte, from, to string) (decimal.Decimal, error)
}

type RequestingCurrencyConverter struct {
	RateRequester RateRequester
	RateDecoder   RateDecoder
}

func (c *RequestingCurrencyConverter) Convert(from, to string, amount decimal.Decimal) (decimal.Decimal, error) {

	response, err := c.RateRequester.Request(from, to)
	if err != nil {
		return decimal.Zero, err
	}
	rate, err := c.RateDecoder.Decode(response, from, to)
	if err != nil {
		return decimal.Zero, err
	}
	if rate.IsZero() {
		return decimal.Zero, fmt.Errorf("Conversion failed: rate is zero")
	}

	return amount.Div(rate), nil
}
