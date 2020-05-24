package dimos

import (
	"bytes"
	"errors"

	"github.com/cbergoon/merkletree"
	"github.com/vmihailenco/msgpack"
	"github.com/wisepythagoras/dimoschain/crypto"
)

// Block represents each individual block in the chain.
type Block struct {
	IDx          int64         `json:"i"`
	MerkleRoot   []byte        `json:"m"`
	Timestamp    int64         `json:"ts"`
	Transactions []Transaction `json:"txs"`
	Hash         []byte        `json:"h"`
	PrevHash     []byte        `json:"ph"`
	Signature    []byte        `json:"s"`
	merkleTree   *merkletree.MerkleTree
}

// AddTransaction adds a transaction to the blockchain.
func (b *Block) AddTransaction(tx *Transaction) bool {
	if tx == nil {
		return false
	}

	// Add the transaction.
	b.Transactions = append(b.Transactions, *tx)

	// Update the Merkle root and the hash.
	b.UpdateHash()

	return true
}

// UpdateHash updates the block's hash
func (b *Block) UpdateHash() error {
	b.ComputeMerkleRoot()

	// Compute the block's hash.
	hash, err := crypto.GetSHA3512Hash(nil)

	if err != nil {
		return err
	}

	b.Hash = hash

	return nil
}

// GetInterface returns the interface.
func (b *Block) GetInterface(includeTx bool) interface{} {
	var returnable interface{}

	if includeTx {
		type BlockRep struct {
			IDx          int64
			MerkleRoot   []byte
			Timestamp    int64
			Transactions []Transaction
			Hash         []byte
			PrevHash     []byte
			Signature    []byte
		}

		returnable = BlockRep{
			IDx:        b.IDx,
			MerkleRoot: b.MerkleRoot,
			Timestamp:  b.Timestamp,
			Hash:       b.Hash,
			PrevHash:   b.PrevHash,
			Signature:  b.Signature,
		}
	} else {
		type BlockRep struct {
			IDx        int64
			MerkleRoot []byte
			Timestamp  int64
			Hash       []byte
			PrevHash   []byte
			Signature  []byte
		}

		returnable = BlockRep{
			IDx:        b.IDx,
			MerkleRoot: b.MerkleRoot,
			Timestamp:  b.Timestamp,
			Hash:       b.Hash,
			PrevHash:   b.PrevHash,
			Signature:  b.Signature,
		}
	}

	return returnable
}

// GetSerialized returns the msgpack version of this block.
func (b *Block) GetSerialized(includeTx bool) ([]byte, error) {
	return msgpack.Marshal(b.GetInterface(includeTx))
}

// ComputeHash computes the hash of the block.
func (b *Block) ComputeHash() ([]byte, error) {
	// Get the msgpack version of the block.
	bin, err := b.GetSerialized(false)

	if err != nil {
		return nil, err
	}

	// Get the hash.
	hash, err := crypto.GetSHA3512Hash(bin)

	if err != nil {
		return nil, err
	}

	b.Hash = hash

	return hash, nil
}

// ComputeMerkleRoot computes the merkle root based on
func (b *Block) ComputeMerkleRoot() ([]byte, error) {
	var list []merkletree.Content

	// Append the transactions to the list of leaves.
	for _, tx := range b.Transactions {
		list = append(list, tx)
	}

	// Create the new Merkle tree.
	tree, err := merkletree.NewTree(list)

	if err != nil {
		return nil, err
	}

	b.merkleTree = tree
	root := tree.MerkleRoot()

	// If there is a merkle root present on the instance and it doesn't match with
	// the computed root, then this means that there is an inconsistency or even
	// attempted forgery.
	if b.MerkleRoot != nil && bytes.Compare(b.MerkleRoot, root) != 0 {
		return nil, errors.New("Invalid root computed")
	}

	b.MerkleRoot = root

	return b.MerkleRoot, nil
}

// VerifyMerkleTree verifies if a transaction is part of the merkle tree.
func (b *Block) VerifyMerkleTreeTx(tx *Transaction) bool {
	// Verify the content in the merkle tree.
	vc, err := b.merkleTree.VerifyContent(tx)

	if err != nil {
		return false
	}

	return vc
}

// BlockFromBytes converts a block directly from the fs into an unserialized Block object.
func BlockFromBytes(b []byte) (*Block, error) {
	if b == nil {
		return nil, errors.New("Nil block bytes")
	}

	block := &Block{}

	// Try to unmarshal the msgpack payload.
	err := msgpack.Unmarshal(b, block)

	return block, err
}
