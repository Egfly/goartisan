package config

import (
	"github.com/Egfly/goartisan/command"
)

var CmdList = map[string]interface{}{
	"version": command.NewVersion(),
}
