package requesting

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type InterestCurrencyConverter struct {
	RateRequester RateRequester
	RateDecoder   RateDecoder
}

func (c *InterestCurrencyConverter) Convert(from, to string, amount decimal.Decimal) (decimal.Decimal, error) {

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

	return amount.Div(rate).Mul(decimal.NewFromFloat(1.05)), nil
}
