package converter

import (
	"github.com/shopspring/decimal"
)

type CurrencyConverter interface {
	Convert(from, to string, amount decimal.Decimal) (decimal.Decimal, error)
}
