package dimos

import (
	"github.com/wisepythagoras/dimoschain/utils"
	badger "github.com/dgraph-io/badger/v2"
)

// Blockchain represents the object that handles the entire blockchain database.
type Blockchain struct {
	Height int64 `json: "h"`
	ID int64 `json: "id"`
	CurrentHash []byte `json: "ch"`
	db badger.DB
}

// SetDB sets the database object onto the current blockchain object.
func (b *Blockchain) SetDB(db badger.DB) {
	b.db = db
}

// LoadChainDB locates and loads the blockchain.
func LoadChainDB() (*Blockchain, error) {
	// Get the chain's directory.
	path, err := utils.GetChainDir(true)

	if err != nil {
		return nil, err
	}

	// Now try to open the database.
	db, err := badger.Open(badger.DefaultOptions(path))

	if err != nil {
		return nil, err
	}

	// Create a new instance of the Blockchain object.
	blockchain := Blockchain{
		Height: 0,
		ID: 0,
		CurrentHash: nil,
	}

	// Set the database onto our new blockchain object.
	blockchain.SetDB(db)

	return &blockchain, nil
}
