package run

import (
	"flag"
	"github.com/Egfly/goartisan/config"
	"io"
	"reflect"
)

type nullIo struct{}

func (ni nullIo) Write(p []byte) (n int, err error) {
	return len(p), nil
}

type runner struct {
	args []string
}

func LoadCommandList() (list map[string]interface{}) {
	list = config.CmdList
	return
}

func Run(w io.Writer, appArgs []string) (string, error) {
	if len(appArgs) == 1 {
		return "no command", nil
	}

	flags := flag.NewFlagSet("goartisan", flag.ContinueOnError)
	flags.SetOutput(nullIo{})
	err := flags.Parse(appArgs[1:])
	if err != nil {
		return "help info", err
	}

	args := flags.Args()
	var cmd interface{}
	cmdList := LoadCommandList()
	for sig, v := range cmdList {
		if sig == args[0] {
			cmd = v
		}
	}
	val := reflect.ValueOf(cmd)
	kd := val.Elem().Kind()
	if kd != reflect.Struct {
		msg := args[0] + "not found"
		return msg, nil
	}

	args = args[1:]
	r := &runner{
		args: args,
	}
	return r.run(val)
}

func (r *runner) run(ref reflect.Value) (string, error) {
	var returns []reflect.Value
	ref.MethodByName("Handle").Call(returns)
	return "", nil
}
