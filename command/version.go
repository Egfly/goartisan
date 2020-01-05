package command

type VersionCommand struct {
	Command
}

func init() {

}

func (v *VersionCommand) Handle() (str string, err error) {
	str = "goartisan version is v1.0.0.0"
	return str, nil
}

func NewVersion() (version *VersionCommand) {
	version = &VersionCommand{}
	version.SetSignature("version")
	version.SetDescription("show go artisan version")

	return
}
