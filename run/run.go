package run

import (
	"flag"
	"github.com/Egfly/goartisan/config"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"reflect"
)

type nullIo struct{}

func (ni nullIo) Write(p []byte) (n int, err error) {
	return len(p), nil
}

type runner struct {
	args []string
}

func BuildPlugin(path, pluginPath string) {
	build := exec.Command("go", "build", "-buildmode=plugin", "-o", pluginPath, path)
	err := build.Run()
	if err != nil {
		panic(err)
	}

	return
}

func LoadCommandList(arg string) (list map[string]interface{}) {
	list = config.CmdList
	// todo 获取文件路径err处理
	dir, _ := filepath.Abs(filepath.Dir(arg))
	dir = dir + "/config/goartisan.go"
	_, err := os.Lstat(dir)

	//判断./config/goartisan.go文件是否存在
	// 存在则将其编译成插件
	if !os.IsNotExist(err) {
		goPath := os.Getenv("GOPATH")
		pluginPath := goPath + "/bin/goartisan.so"
		BuildPlugin(dir, pluginPath)
		// todo plugin
		p, err := plugin.Open(pluginPath)
		if err != nil {
			panic(err)
		}
		cl, err := p.Lookup("CommandList")
		if err != nil {
			panic(err)
		}
		res := cl.(*map[string]interface{})
		for k, v := range *res {
			list[k] = v
		}

	}
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
	cmdList := LoadCommandList(appArgs[0])
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

	result := ref.MethodByName("Handle").Call(nil)

	return result[0].String(), nil
}
