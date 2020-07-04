package utils

// Name is the name of the blockchain.
const (
	Name = "Dimosthenes"

	// Version is the current version of the blockchain.
	Version = "0.1.1"

	// DimosDir is the directory that the blockchain will live in.
	DimosDir = ".dimos"

	// ChainDir is the directory where the blockchain database will live in.
	ChainDir = "chain"

	// Genesis is the name of the file that will contain the genesis hash.
	Genesis = "genesis"

	// CurrentHash is the file that contains the current hash.
	CurrentHash = "current_hash"

	// MaxUnitSupply defines the maximum unit supply, which is 1 qadrillion.
	MaxUnitSupply = 1e15

	// UnitsInCoin The amount of units in a coin are 100 million. As a result, we can have 10 million coins.
	UnitsInCoin = 1e8
)
