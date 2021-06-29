package application

import (
	"fmt"
	"io"
)

type Application interface {
	Run(args []string, stdout, stderr io.Writer) error
}

func RunApplication(app Application, args []string, stdout, stderr io.Writer) error {

	err := app.Run(args, stdout, stderr)
	if err != nil {
		fmt.Fprintf(stderr, "%v\n", err)
		return err
	}
	return nil
}
