package run

import (
	"context"
	"flag"
	"fmt"
	"io"
)

type nullIo struct{}

func (ni nullIo) Write(p []byte) (n int, err error) {
	return len(p), nil
}

type runner struct {
	ctx *context.Context
}

func Run(w io.Writer, appArgs []string) (string, error) {
	r := &runner{}
	return r.run(w, appArgs)
}

func (r *runner) run(w io.Writer, appArgs []string) (string, error) {
	if len(appArgs) == 1 {
		return "no command", nil
	}

	flags := flag.NewFlagSet("goartisan", flag.ContinueOnError)
	//licenses := flags.Bool("govendor-licenses", false, "show govendor's licenses")
	version := flags.Bool("version", false, "show govendor version")
	flags.SetOutput(nullIo{})
	err := flags.Parse(appArgs[1:])
	if err != nil {
		return "help info", err
	}

	if version != nil {
		return "v1.0.0.0", nil
	}

	args := flags.Args()

	cmd := args[0]

	switch cmd {
	case "version":
		return "v1.0.0.0", nil
	default:
		return "", fmt.Errorf("Unknown command %q", cmd)
	}
}
