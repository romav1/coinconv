package requesting

type RateRequester interface {
	Request(from, to string) ([]byte, error)
}

type RateDecoder interface {
	Decode(raw []byte, from, to string) (float64, error)
}

type RequestingCurrencyConverter struct {
	RateRequester RateRequester
	RateDecoder   RateDecoder
}

func (c *RequestingCurrencyConverter) Convert(from, to string, amount float64) (float64, error) {

	response, err := c.RateRequester.Request(from, to)
	if err != nil {
		return 0, err
	}
	rate, err := c.RateDecoder.Decode(response, from, to)
	if err != nil {
		return 0, err
	}
	if rate == 0 {
		return 0, nil
	}

	return amount / rate, nil
}
