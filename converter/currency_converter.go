package converter

type CurrencyConverter interface {
	Convert(from, to string, amount float64) (float64, error)
}
