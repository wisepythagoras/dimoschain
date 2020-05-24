package dimos

import (
	"errors"

	badger "github.com/dgraph-io/badger"
	"github.com/wisepythagoras/dimoschain/utils"
)

// Blockchain represents the object that handles the entire blockchain database.
type Blockchain struct {
	Height      int64  `json:"h"`
	ID          int64  `json:"id"`
	CurrentHash []byte `json:"ch"`
	genesisHash []byte
	db          *badger.DB
}

// GetDB returns the genesis hash.
func (b *Blockchain) GetDB() []byte {
	return b.genesisHash
}

// GetCurrentBlock gets the current block.
func (b *Blockchain) GetCurrentBlock() (*Block, error) {
	return b.GetBlock(b.CurrentHash)
}

// GetBlock get's a block by its hash.
func (b *Blockchain) GetBlock(hash []byte) (*Block, error) {
	if hash == nil {
		return nil, errors.New("Nil hash")
	}

	// Create a new transaction.
	txn := b.db.NewTransaction(true)

	// Get the item of the entry with the hash as the key.
	item, err := txn.Get(hash)

	if err != nil {
		return nil, err
	}

	// Get the vaue from the item.
	value, err := item.ValueCopy(nil)

	if err != nil {
		return nil, err
	}

	// Parse the block.
	return BlockFromBytes(value)
}

// CreateBlock adds a block to the chain.
func (b *Blockchain) AddBlock(block *Block) (bool, error) {
	if block == nil {
		return false, errors.New("Invalid block")
	}

	isGenesisBlock := block.IDx == 1

	// If the id is 1, this means that we are trying to add the genesis block, so we
	// don't need a current or genesis hash.
	if !isGenesisBlock && (b.CurrentHash == nil || b.genesisHash == nil) {
		return false, errors.New("The blockchain has not been initialized")
	}

	if !isGenesisBlock {
		currentBlock, err := b.GetCurrentBlock()

		if err != nil {
			return false, err
		}

		block.IDx = currentBlock.IDx + 1
		block.PrevHash = currentBlock.Hash
		block.ComputeHash()
	}

	// Create a new transaction.
	txn := b.db.NewTransaction(true)
	defer txn.Discard()

	// Get the serialized block.
	serialized, err := block.GetSerialized(true)

	if err != nil {
		return false, err
	}

	// Set the block onto the database.
	if err = txn.Set(block.Hash, serialized); err != nil {
		return false, err
	}

	// Commit the changes to the database.
	if err = txn.Commit(); err != nil {
		return false, err
	}

	// Write the current hash into the current hash file on the disk.
	utils.WriteCurrentHash(block.Hash)

	return true, nil
}

// CreateChainInstance creates a new instance of the blockchain object.
func CreateChainInstance(genesisHash []byte, currentHash []byte) (*Blockchain, error) {
	// Get the chain's directory.
	path, err := utils.GetChainDir(true)

	if err != nil {
		return nil, err
	}

	// Now try to open the database.
	db, err := badger.Open(badger.DefaultOptions(path + "/" + utils.CHAIN_DIR))

	if err != nil {
		return nil, err
	}

	// Create a new instance of the Blockchain object.
	blockchain := Blockchain{
		Height:      0,
		ID:          0,
		CurrentHash: currentHash,
		genesisHash: genesisHash,
		db:          db,
	}

	return &blockchain, nil
}

// InitChainDB locates and loads the blockchain.
func InitChainDB() (*Blockchain, error) {
	// Get the genesis block. If it doesn't exist, then the databse hasn't been
	// initialized.
	genesisHash, err := utils.GetGenesisHash()

	if err != nil {
		return nil, err
	}

	// Get the current hash.
	currentHash, err := utils.GetCurrentHash()

	if err != nil {
		return nil, err
	}

	// Create a new instance of the blockchain object and return.
	return CreateChainInstance(genesisHash, currentHash)
}
