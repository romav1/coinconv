package coinmarketcap

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type CoinmarketcapQuotesRateRequester struct {
	ApiKey  string
	ApiHost string
}

func (r *CoinmarketcapQuotesRateRequester) Request(from, to string) ([]byte, error) {

	ApiProtocol := "https://"
	ApiUrl := "/v1/cryptocurrency/quotes/latest"

	client := &http.Client{}
	req, err := http.NewRequest("GET", ApiProtocol+r.ApiHost+ApiUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("Request creation error: %v", err)
	}

	q := url.Values{}
	q.Add("convert", from)
	q.Add("symbol", to)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", r.ApiKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request to server: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad response status: %v", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	return respBody, err
}
