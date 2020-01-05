package config

import "github.com/Egfly/goartisan/command"

var CmdList = map[string]interface{}{
	"c": command.VersionCommand{},
}
