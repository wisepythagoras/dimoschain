package main

import (
	"log"

	"github.com/wisepythagoras/dimoschain/dimos"
)

func main() {
	// Load the database.
	blockchain, err := dimos.LoadChainDB()

	if err != nil {
		log.Fatal(err)
	}

	// Create a test transaction.
	tx := dimos.Transaction{
		Hash:      nil,
		Amount:    0.001,
		From:      []byte("test1"),
		To:        []byte("test2"),
		Signature: []byte("test1signature"),
	}

	block := dimos.Block{}
	block.AddTransaction(&tx)

	blockchain.AddBlock(&block)
}
