package console

type CmdInterface interface {
	CmdConfig() (trigger string, description string)
	CmdExecute() bool
}
