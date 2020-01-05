package command

type Handle interface {
	Handle() (string, error)
	GetSignature() string
	SetSignature(signature string)
	GetDescription() string
	SetDescription(des string)
}
