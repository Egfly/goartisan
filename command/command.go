package command

type Command struct {
	Signature   string
	Description string
}

func (cmd *Command) GetSignature() string {
	return cmd.Signature
}

func (cmd *Command) SetSignature(signature string) {
	cmd.Signature = signature
}

func (cmd *Command) GetDescription() string {
	return cmd.Signature
}

func (cmd *Command) SetDescription(des string) {
	cmd.Description = des
}
