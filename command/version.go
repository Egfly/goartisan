package command

type VersionCommand struct {
	Command
}

func init() {
	version := VersionCommand{}
	version.SetSignature("version")
	version.SetDescription("show go artisan version")
}

func (v *VersionCommand) Handle() (str string, err error) {
	str = "goartisan version is v1.0.0.0"
	return str, nil
}
