package requesting

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/romav1/coinconv/converter"
)

type CachingCurrencyConverter struct {
	Cache     Cache
	Converter converter.CurrencyConverter
}

var _ converter.CurrencyConverter = (*CachingCurrencyConverter)(nil)

func (c *CachingCurrencyConverter) Convert(from, to string, amount decimal.Decimal) (decimal.Decimal, error) {

	price, err := c.Cache.Get(from + to)
	if err == nil && !price.IsZero() {
		return amount.Div(price), nil
	}

	res, err := c.Converter.Convert(from, to, amount)

	if err != nil {
		return decimal.Zero, fmt.Errorf("conversion error: %v", err)
	}

	err = c.Cache.Set(from+to, res.Mul(amount))

	if err != nil {

		return res, fmt.Errorf("cache storing error %v\n", err)
	}
	return res, nil

}
