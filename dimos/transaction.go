package dimos

import (
	"bytes"
	"encoding/json"

	"github.com/cbergoon/merkletree"
	"github.com/wisepythagoras/dimoschain/crypto"
	"github.com/wisepythagoras/dimoschain/utils"
)

// Transaction represents a single transaction from and to another wallet in the
// dimosthenes network.
type Transaction struct {
	Hash      []byte  `json:"h"`
	Amount    float64 `json:"a"`
	From      []byte  `json:"f"`
	To        []byte  `json:"t"`
	Signature []byte  `json:"s"`
}

// CalculateHash calculates the hash of this transaction.
func (tx Transaction) CalculateHash() ([]byte, error) {
	var hashFormat []byte
	hashFormat, err := utils.Float64ToByte(tx.Amount)

	if err != nil {
		return nil, err
	}

	hashFormat = append(hashFormat, tx.From...)
	hashFormat = append(hashFormat, tx.To...)
	hashFormat = append(hashFormat, tx.Signature...)

	// Make the JSON format of the transaction we're going to hash.
	jsonTx, err := json.Marshal(hashFormat)

	if err != nil {
		return nil, err
	}

	// Calculate the hash.
	hash, err := crypto.GetSHA3512Hash(jsonTx)

	if err != nil {
		return nil, err
	}

	return hash, nil
}

// Equals checks two transactions for equality.
func (tx Transaction) Equals(otherTx merkletree.Content) (bool, error) {
	return bytes.Compare(tx.Hash, otherTx.(Transaction).Hash) == 1, nil
}
