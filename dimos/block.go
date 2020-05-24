package dimos

import (
	"bytes"
	"errors"

	"github.com/cbergoon/merkletree"
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
	if b.MerkleRoot != nil && bytes.Compare(b.MerkleRoot, root) == 0 {
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
