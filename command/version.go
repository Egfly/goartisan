package command

import "github.com/Egfly/goartisan/config"

type VersionCommand struct {
	Command
}

func init() {
	version := VersionCommand{}
	version.SetSignature("version")
	version.SetDescription("show go artisan version")
}

func (v *VersionCommand) Handle() (str string, err error) {
	str = "goartisan version is" + config.Version
	return str, nil
}
