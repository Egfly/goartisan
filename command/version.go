package command

import "goartisan/config"

type VersionCommand struct {
	Command
}

func (v *VersionCommand) Handle() (str string, err error) {
	str = "goartisan version is" + config.Version
	return str, nil
}
