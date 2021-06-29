package coinmarketcap

import (
	"fmt"
	"testing"
)

func TestCoinmarketcapQuotesRateRequester(t *testing.T) {

	type Testcase struct {
		ApiKey     string
		ApiHost    string
		From       string
		To         string
		Result     string
		ShouldPass bool
	}

	testcases := []Testcase{
		{
			ApiKey:  "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c",
			ApiHost: "sandbox-api.coinmarketcap.com",
			From:    "BTC",
			To:      "USD",
			// we can't test for result here as the sandbox result is random
			ShouldPass: true,
		},
		{
			ApiKey:  "xxxxx-b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c",
			ApiHost: "pro-api.coinmarketcap.com",
			From:    "BTC",
			To:      "USD",
			// we can't test for result here as the sandbox result is random
			ShouldPass: false,
		},
		{
			ApiKey:  "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c",
			ApiHost: "xxxx.sandbox-api.coinmarketcap.com",
			From:    "BTC",
			To:      "USD",
			// we can't test for result here as the sandbox result is random
			ShouldPass: false,
		},
	}

	for i, testcase := range testcases {
		requester := CoinmarketcapQuotesRateRequester{ApiKey: testcase.ApiKey, ApiHost: testcase.ApiHost}

		result, err := requester.Request(testcase.From, testcase.To)

		fmt.Println(string(result), err)

		if err != nil && testcase.ShouldPass {
			t.Fatalf("Testcase %v should pass, but failed: %v", testcase, err)
		}
		if err == nil && !testcase.ShouldPass {
			t.Fatalf("Testcase %v should fail, but passed: %v", testcase, string(result))
		}
		fmt.Printf("Requester PASS %v\n", i)

	}
}
