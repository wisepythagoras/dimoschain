package dimos

import (
	_ "github.com/cbergoon/merkletree"
)

type Block struct {
	MerkleRoot []byte `json: "m"`
	Transactions []Transaction `json: "txs"`
	Hash []byte `json: "h"`
	Signature []byte `json: "s"`
}
