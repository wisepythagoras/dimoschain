package main

import (
	"encoding/hex"
	"log"
	"time"

	"github.com/wisepythagoras/dimoschain/dimos"
	"github.com/wisepythagoras/dimoschain/utils"
)

func main() {
	// Load the database.
	blockchain, err := dimos.CreateChainInstance(nil, nil)

	if err != nil {
		log.Fatal(err)
		return
	}

	// If there is a current has, the chain has already been instanciated and we should not move on.
	currentHash, err := utils.GetCurrentHash()

	if currentHash != nil || err == nil {
		log.Fatal("The blockchain has already been instanciated")
		return
	}

	// Dummy transaction.
	tx := dimos.Transaction{
		Amount:    0,
		From:      []byte("0"),
		To:        []byte("0"),
		Signature: []byte("0"),
	}
	tx.Hash, err = tx.CalculateHash()

	// This is the time of genesis.
	date, _ := time.Parse(time.RFC3339, "2018-04-05T19:24:45Z")

	// Create the genesis block.
	genesisBlock := dimos.Block{
		IDx:          1,
		Timestamp:    date.Unix(),
		Transactions: []dimos.Transaction{},
		PrevHash:     []byte("0"),
		Signature:    []byte("0"),
	}

	// Add a transaction. The merkle root is updated every time we append a new
	// transaction.
	genesisBlock.AddTransaction(&tx)

	// The merkle root is updated every time we append a new transaction, but we get
	// it here, so that we can catch any error.
	_, err = genesisBlock.ComputeMerkleRoot(false)

	log.Println("Merkle Root:", hex.EncodeToString(genesisBlock.MerkleRoot), err)

	// Compute the hash of the block.
	hash, err := genesisBlock.ComputeHash(false)

	log.Println("Genesis Hash:", hex.EncodeToString(hash), err)

	// Write the genesis hash to the disk.
	_ = utils.WriteGenesisHash(hash)

	// Add the block.
	if _, err = blockchain.AddBlock(&genesisBlock); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Created genesis")
	}
}
