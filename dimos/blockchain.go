package dimos

import (
	badger "github.com/dgraph-io/badger"
	"github.com/wisepythagoras/dimoschain/utils"
)

// Blockchain represents the object that handles the entire blockchain database.
type Blockchain struct {
	Height      int64  `json:"h"`
	ID          int64  `json:"id"`
	CurrentHash []byte `json:"ch"`
	db          *badger.DB
}

// SetDB sets the database object onto the current blockchain object.
func (b *Blockchain) SetDB(db *badger.DB) {
	b.db = db
}

// CreateBlock adds a block to the chain.
func (b *Blockchain) AddBlock(block *Block) (bool, error) {
	return true, nil
}

// LoadChainDB locates and loads the blockchain.
func LoadChainDB() (*Blockchain, error) {
	// Get the chain's directory.
	path, err := utils.GetChainDir(true)

	if err != nil {
		return nil, err
	}

	// Now try to open the database.
	db, err := badger.Open(badger.DefaultOptions(path + "/chain"))

	if err != nil {
		return nil, err
	}

	// Create a new instance of the Blockchain object.
	blockchain := Blockchain{
		Height:      0,
		ID:          0,
		CurrentHash: nil,
	}

	// Set the database onto our new blockchain object.
	blockchain.SetDB(db)
	defer db.Close()

	return &blockchain, nil
}
