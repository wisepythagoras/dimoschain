package dimos

import (
	"bytes"
	"errors"
	"github.com/cbergoon/merkletree"
)

// Block represents each individual block in the chain.
type Block struct {
	MerkleRoot []byte `json: "m"`
	merkleTree *merkletree.MerkleTree
	Transactions []Transaction `json: "txs"`
	Hash []byte `json: "h"`
	Signature []byte `json: "s"`
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
