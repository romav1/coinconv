package requesting

import (
	"github.com/shopspring/decimal"
)

type FakeCache struct {
}

func (c *FakeCache) Get(Pair string) (decimal.Decimal, error) {

	return decimal.Zero, nil
}

func (c *FakeCache) Set(Pair string, amount decimal.Decimal) error {

	return nil
}
