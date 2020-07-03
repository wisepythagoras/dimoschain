package proto

const (
	// Transaction is for creating a new transaction for sending tokens to another address
	// or anything else that the protocol allows.
	Transaction = iota

	// Update is for receiving any blocks that follow the current index of the client.
	Update = iota
)
