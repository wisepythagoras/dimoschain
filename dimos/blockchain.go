package dimos

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"

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

// IsChainValid checks if the blockchain is consistent and that all blocks are
// valid. Every block in the chain needs to have the correct index (IDx), hash
// and previous hash.
func (b *Blockchain) IsChainValid() (bool, error) {
	// First check if the blockchain has been instanciated.
	if b.db == nil {
		return false, errors.New("No instance of the blockchain")
	}

	// Then, the next step would be to get the current block. From there on we
	// will go on to the previous blocks until we reach the genesis block.
	block, err := b.GetCurrentBlock()

	if err != nil {
		return false, err
	}

	// We'll use this as our next block reference.
	nextBlock := block

	// Now we loop. This is probably not efficient, and it will be rewritten in the
	// future, but for now it stays put.
	for block, err = b.GetBlock(block.PrevHash); err == nil && block != nil &&
		bytes.Compare(block.PrevHash, []byte("0")) == 0; {
		// Here, technically we will, at some point, reach the genesis block. This
		// means that the loop will exit when the genesis block or an error is reached.

		// Just validate the block.
		if _, err = b.ValidateBlock(nextBlock, block); err != nil {
			return false, err
		}

		nextBlock = block
	}

	if err != nil {
		return false, err
	}

	return true, nil
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

// ValidateBlock validates that a block conforms with the rules of our blockchain. Which means
// that the prevHash of our new block needs to match the hash of the prevBlock. Also, the hash
// of the block, as well as the merkle root, need to check out.
func (b *Blockchain) ValidateBlock(block *Block, prevBlock *Block) (bool, error) {
	// First compare the blocks.
	if bytes.Compare(block.PrevHash, prevBlock.Hash) != 0 {
		return false, errors.New("The prevHash doesn't match the hash of given prevBlock")
	}

	// The IDx of our block needs to be an increment of 1 above the previous block.
	if prevBlock.IDx != block.IDx-1 {
		str := fmt.Sprintf("Invalid IDx found at block %s", hex.EncodeToString(block.Hash))
		return false, errors.New(str)
	}

	// Check the signature here.

	// Now verify the merkle root.
	merkleRoot, err := block.ComputeMerkleRoot()

	if err != nil || bytes.Compare(merkleRoot, block.MerkleRoot) == 0 {
		if err == nil {
			str := fmt.Sprintf("Invalid merkle root at block %s", hex.EncodeToString(block.Hash))
			err = errors.New(str)
		}

		return false, err
	}

	// Lastly we check the hash of the block.
	hash, err := block.ComputeHash()

	if err != nil || bytes.Compare(hash, block.Hash) == 0 {
		if err == nil {
			str := fmt.Sprintf("Invalid block hash at %s", hex.EncodeToString(block.Hash))
			err = errors.New(str)
		}

		return false, err
	}

	return true, nil
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
