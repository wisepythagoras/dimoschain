package dimos

import (
	"bytes"
	"encoding/hex"
	"fmt"

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

	// Calculate the hash.
	hash, err := crypto.GetSHA3384Hash(hashFormat)

	if err != nil {
		return nil, err
	}

	tx.Hash = hash

	return hash, nil
}

// Equals checks two transactions for equality.
func (tx Transaction) Equals(otherTx merkletree.Content) (bool, error) {
	return bytes.Compare(tx.Hash, otherTx.(Transaction).Hash) == 1, nil
}

// String returns the string representation of the transaction.
func (tx Transaction) String() string {
	return "Tx: " + hex.EncodeToString(tx.Hash) + "\n" +
		" Amount: " + fmt.Sprintf("%.10f", tx.Amount) + "\n" +
		" From: " + string(tx.From) + "\n" +
		" To: " + string(tx.To) + "\n" +
		" Signature: " + hex.EncodeToString(tx.Signature)
}
