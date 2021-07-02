package coinmarketcap

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestCoinmarketcapQuotesRateDecoder(t *testing.T) {

	type Testcase struct {
		Input      string
		From       string
		To         string
		Output     string
		ShouldPass bool
	}

	testcases := []Testcase{
		{
			Input:      `{"status":{"timestamp":"2021-06-28T21:28:26.824Z","error_code":0,"error_message":null,"elapsed":1,"credit_count":1,"notice":null},"data":{"BTC":{"id":6887,"name":"zcypjrnvmj","symbol":"BTC","slug":"w3kefp8rm8r","is_active":6708,"is_fiat":null,"circulating_supply":8193,"total_supply":6376,"max_supply":7698,"date_added":"2021-06-28T21:28:26.824Z","num_market_pairs":9646,"cmc_rank":2799,"last_updated":"2021-06-28T21:28:26.824Z","tags":["ar7x0ajjljq","fy2y8x6sk7","l0w6wlia5zo","s4omwkaxlhk","uaczbi5bqwl","kttbj83qqo","pz6sdzcr4jg","fnz516j9xlk","dtwocii0qe","9kjnou61ki"],"platform":null,"quote":{"USD":{"price":0.5177489085472153,"volume_24h":0.7057734076665365,"percent_change_1h":0.9123102248041741,"percent_change_24h":0.645121093850662,"percent_change_7d":0.7670986097277475,"percent_change_30d":0.5552903336384081,"market_cap":0.09933535156822448,"last_updated":"2021-06-28T21:28:26.824Z"}}}}}`,
			From:       "USD",
			To:         "BTC",
			Output:     "0.5177489085472153",
			ShouldPass: true,
		},
		{
			Input:      `{"status":{"timestamp":"2021-06-28T21:28:26.824Z","error_code":0,"error_message":null,"elapsed":1,"credit_count":1,"notice":null},"data":{"BTC":{"id":6887,"name":"zcypjrnvmj","symbol":"BTC","slug":"w3kefp8rm8r","is_active":6708,"is_fiat":null,"circulating_supply":8193,"total_supply":6376,"max_supply":7698,"date_added":"2021-06-28T21:28:26.824Z","num_market_pairs":9646,"cmc_rank":2799,"last_updated":"2021-06-28T21:28:26.824Z","tags":["ar7x0ajjljq","fy2y8x6sk7","l0w6wlia5zo","s4omwkaxlhk","uaczbi5bqwl","kttbj83qqo","pz6sdzcr4jg","fnz516j9xlk","dtwocii0qe","9kjnou61ki"],"platform":null,"quote":{"USD":{"price":0.5177489085472153,"volume_24h":0.7057734076665365,"percent_change_1h":0.9123102248041741,"percent_change_24h":0.645121093850662,"percent_change_7d":0.7670986097277475,"percent_change_30d":0.5552903336384081,"market_cap":0.09933535156822448,"last_updated":"2021-06-28T21:28:26.824Z"}}}}}`,
			From:       "USD",
			To:         "ETH",
			Output:     "0",
			ShouldPass: false,
		},
		{
			Input:      `{"status":{"timestamp":"2021-06-28T21:28:26.824Z","error_code":0,"error_message":null,"elapsed":1,"credit_count":1,"notice":null},"data":{"BTC":{"id":6887,"name":"zcypjrnvmj","symbol":"BTC","slug":"w3kefp8rm8r","is_active":6708,"is_fiat":null,"circulating_supply":8193,"total_supply":6376,"max_supply":7698,"date_added":"2021-06-28T21:28:26.824Z","num_market_pairs":9646,"cmc_rank":2799,"last_updated":"2021-06-28T21:28:26.824Z","tags":["ar7x0ajjljq","fy2y8x6sk7","l0w6wlia5zo","s4omwkaxlhk","uaczbi5bqwl","kttbj83qqo","pz6sdzcr4jg","fnz516j9xlk","dtwocii0qe","9kjnou61ki"],"platform":null,"quote":{"USD":{"price":0.5177489085472153,"volume_24h":0.7057734076665365,"percent_change_1h":0.9123102248041741,"percent_change_24h":0.645121093850662,"percent_change_7d":0.7670986097277475,"percent_change_30d":0.5552903336384081,"market_cap":0.09933535156822448,"last_updated":"2021-06-28T21:28:26.824Z"}}}}}`,
			From:       "XMR",
			To:         "BTC",
			Output:     "0",
			ShouldPass: false,
		},
		{
			Input:      `{"status":{"timestamp":"2021-06-28T21:28:26.824Z","error_code":1,"error_message":null,"elapsed":1,"credit_count":1,"notice":null},"data":{"BTC":{"id":6887,"name":"zcypjrnvmj","symbol":"BTC","slug":"w3kefp8rm8r","is_active":6708,"is_fiat":null,"circulating_supply":8193,"total_supply":6376,"max_supply":7698,"date_added":"2021-06-28T21:28:26.824Z","num_market_pairs":9646,"cmc_rank":2799,"last_updated":"2021-06-28T21:28:26.824Z","tags":["ar7x0ajjljq","fy2y8x6sk7","l0w6wlia5zo","s4omwkaxlhk","uaczbi5bqwl","kttbj83qqo","pz6sdzcr4jg","fnz516j9xlk","dtwocii0qe","9kjnou61ki"],"platform":null,"quote":{"USD":{"price":0.5177489085472153,"volume_24h":0.7057734076665365,"percent_change_1h":0.9123102248041741,"percent_change_24h":0.645121093850662,"percent_change_7d":0.7670986097277475,"percent_change_30d":0.5552903336384081,"market_cap":0.09933535156822448,"last_updated":"2021-06-28T21:28:26.824Z"}}}}}`,
			From:       "USD",
			To:         "BTC",
			Output:     "0",
			ShouldPass: false,
		},
		{
			Input:      `{"status":{"}}}`,
			From:       "USD",
			To:         "BTC",
			Output:     "0",
			ShouldPass: false,
		},
	}

	var decoder CoinmarketcapQuotesRateDecoder
	for i, testcase := range testcases {
		output, err := decoder.Decode([]byte(testcase.Input), testcase.From, testcase.To)
		outputFloat := output.String()

		if testcase.ShouldPass && err != nil {
			t.Fatalf("Testcase %v should pass, but failed: (%v, %v)", testcase, output, err)
		}
		if !testcase.ShouldPass && err == nil {
			t.Fatalf("Testcase %v should fail, but passed (%v, %v)", testcase, output, err)
		}

		if outputFloat != testcase.Output {
			t.Fatalf("Testcase %v should return %v, but returned %v", testcase, testcase.Output, output)
		}
		fmt.Printf("Decoder PASS %v\n", i)
	}
}
