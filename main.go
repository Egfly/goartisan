package main

import (
	"bytes"
	"flag"
	"fmt"
	"goartisan/run"
	"io"
	"os"
	"strings"
)

func main() {
	allArgs := os.Args

	if allArgs[len(allArgs)-1] == "-" {
		stdin := &bytes.Buffer{}
		if _, err := io.Copy(stdin, os.Stdin); err == nil {
			stdinArgs := strings.Fields(stdin.String())
			allArgs = append(allArgs[:len(allArgs)-1], stdinArgs...)
		}
	}

	msg, err := run.Run(os.Stdout, allArgs)
	if err == flag.ErrHelp {
		err = nil
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
	}
	msgText := msg
	if len(msgText) > 0 {
		fmt.Fprint(os.Stderr, msgText)
	}
	if err != nil {
		os.Exit(2)
	}
}
