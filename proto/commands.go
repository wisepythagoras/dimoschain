package proto

const (
	// CmdTxSend is for creating a new transaction for sending tokens to another address
	// or anything else that the protocol allows.
	CmdTxSend = iota

	// CmdUpdate is for receiving any blocks that follow the current index of the client.
	CmdUpdate = iota

	// CmdExit signals the termination of a client-session.
	CmdExit = iota
)
