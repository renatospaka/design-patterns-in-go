package command
type MoneyTransfering interface {
	Call()
	// Undo()
	// Succeded() bool
	// SetSucceded(value bool)
}

type MoneyTransferCommand struct {
	CompositeBankCommand
	from, to *BankAccount
	amount   int
}

func NewMoneyTransferCommand(from *BankAccount, to *BankAccount, amount int) *MoneyTransferCommand {
	mTrasnfer := &MoneyTransferCommand{
		from:                 from,
		to:                   to,
		amount:               amount,
	}
	mTrasnfer.commands = append(mTrasnfer.commands, NewBankAccountCommand(from, Withdraw, amount))
	mTrasnfer.commands = append(mTrasnfer.commands, NewBankAccountCommand(to, Deposit, amount))

	return mTrasnfer
}

func (m *MoneyTransferCommand) Call() {
	ok := true
	for _, cmd := range m.commands {
		if ok {
			cmd.Call()
			ok = cmd.Succeded()
		} else {
			cmd.SetSucceded(false)
		}
	}
}
