package command

type CompositeBankCommand struct {
	commands []Command
}

type CompositeBanking interface {
	Call()
	Undo()
	Succeded() bool
	SetSucceded(value bool)
}

func (c *CompositeBankCommand) Call() {
	for _, cmd := range c.commands {
		cmd.Call()
	}
}

func (c *CompositeBankCommand) Undo() {
	// need to be execute from last to first
	for id := range c.commands {
		c.commands[len(c.commands) - id - 1].Undo()
	}
}

func (c *CompositeBankCommand) Succeded() bool {
	for _, cmd := range c.commands {
		if !cmd.Succeded() {
			return false
		}
	}
	return true
}

func (c *CompositeBankCommand) SetSucceded(value bool) {
	for _, cmd := range c.commands {
		cmd.SetSucceded(value)
	}
}
