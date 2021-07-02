package requesting

import (
	"github.com/shopspring/decimal"
)

type Cache interface {
	Get(Pair string) (decimal.Decimal, error)
	Set(Pair string, amount decimal.Decimal) error
}
