package core

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/cbergoon/merkletree"
	"github.com/wisepythagoras/dimoschain/crypto"
	"github.com/wisepythagoras/dimoschain/utils"
)

type TxType uint8

const (
	TxEmpty TxType = iota
	TxTransfer
	TxContractCreate
	TxContractCall
)

// Transaction represents a single transaction from and to another wallet in the
// dimosthenes network.
type Transaction struct {
	Hash      []byte `json:"h"`
	Type      TxType `json:"tt"`
	Amount    uint64 `json:"a"`
	From      []byte `json:"f"`
	To        []byte `json:"d"`
	Nonce     uint64 `json:"n"`
	Timestamp int64  `json:"t"`
	Payload   []byte `json:"p"`
	Signature []byte `json:"s"` // TODO: Use Dilithium/SPHINCS+ for quantum-safe crypto.
}

// CalculateHash calculates the hash of this transaction.
func (tx *Transaction) CalculateHash() ([]byte, error) {
	var hashFormat []byte
	hashFormat = utils.UInt64ToBytes(tx.Amount)

	hashFormat = append(hashFormat, tx.From...)
	hashFormat = append(hashFormat, tx.To...)
	hashFormat = append(hashFormat, utils.UInt64ToBytes(tx.Amount)...)
	hashFormat = append(hashFormat, utils.UInt64ToBytes(tx.Nonce)...)
	hashFormat = append(hashFormat, utils.UInt64ToBytes(uint64(tx.Type))...)
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
	return bytes.Compare(tx.Hash, otherTx.(*Transaction).Hash) == 1, nil
}

func (tx *Transaction) f() {}

// String returns the string representation of the transaction.
func (tx Transaction) String() string {
	return fmt.Sprintf(
		"Tx: %s\n Amount: %s\n From: %s\n To: %s\n Nonce: %d\n Ts: %s\n Type: %d\n Signature: %s\n",
		hex.EncodeToString(tx.Hash),
		fmt.Sprintf("%.10f", (float64(tx.Amount)/utils.UnitsInCoin)),
		string(tx.From),
		string(tx.To),
		tx.Nonce,
		time.UnixMilli(tx.Timestamp),
		tx.Type,
		hex.EncodeToString(tx.Signature),
	)
}
