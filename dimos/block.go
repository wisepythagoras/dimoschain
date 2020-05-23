package dimos

import (
	_ "github.com/cbergoon/merkletree"
)

type Block struct {
	MerkleRoot []byte
	Transactions []Transaction
}
