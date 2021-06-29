package main

import (
	"os"

	"github.com/romav1/coinconv/application"
	"github.com/romav1/coinconv/converter/requesting"
	"github.com/romav1/coinconv/converter/requesting/coinmarketcap"
)

func main() {

	converter := &requesting.RequestingCurrencyConverter{
		RateRequester: &coinmarketcap.CoinmarketcapQuotesRateRequester{
			ApiKey:  "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c",
			ApiHost: "sandbox-api.coinmarketcap.com",
		},
		RateDecoder: &coinmarketcap.CoinmarketcapQuotesRateDecoder{},
	}
	capp := &application.ConvertApplication{Converter: converter}

	err := application.RunApplication(capp, os.Args, os.Stdout, os.Stderr)

	if err != nil {
		os.Exit(1)
	}
}
