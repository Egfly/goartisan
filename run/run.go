package run

import (
	"context"
	"flag"
	"fmt"
	"goartisan/config"
	"io"
)

type nullIo struct{}

func (ni nullIo) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func init() {
	fmt.Println("this is init func")
}

type runner struct {
	ctx *context.Context
}

func Run(w io.Writer, appArgs []string) (string, error) {
	if len(appArgs) == 1 {
		return "no command", nil
	}

	flags := flag.NewFlagSet("goartisan", flag.ContinueOnError)
	//licenses := flags.Bool("govendor-licenses", false, "show govendor's licenses")
	version := flags.Bool("version", false, "show goartisan version")
	flags.SetOutput(nullIo{})
	err := flags.Parse(appArgs[1:])
	if err != nil {
		return "help info", err
	}

	if version != nil {
		return config.Version, nil
	}

	args := flags.Args()
	cmd := args[0]
	r := &runner{}
	//switch cmd {
	//case "version":
	//	return "v1.0.0.0", nil
	//default:
	//	return "", fmt.Errorf("Unknown command %q", cmd)
	//}

	return r.run(cmd, args[1:])
}

func (r *runner) run(signature string, args []string) (string, error) {

}
