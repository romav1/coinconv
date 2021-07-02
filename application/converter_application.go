package application

import (
	"fmt"
	"io"

	"github.com/romav1/coinconv/converter"

	"github.com/shopspring/decimal"
)

type ConvertApplication struct {
	Converter converter.CurrencyConverter
}

func (capp *ConvertApplication) Run(args []string, stdout, stderr io.Writer) error {

	if len(args) == 0 {
		return fmt.Errorf("usage: app <amount> <from> <to>\n")
	}
	if len(args) < 4 {
		return fmt.Errorf("usage: %s <amount> <from> <to>\n", args[0])
	}
	amount, err := decimal.NewFromString(args[1])
	if err != nil {
		return fmt.Errorf("amount parsing error: %v", err)
	}

	result, err := capp.Converter.Convert(args[2], args[3], amount)
	if err != nil {
		return fmt.Errorf("conversion error: %v", err)
	}
	fmt.Fprintf(stdout, "%v\n", result)
	return nil
}
