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
	}

	// Create the genesis block.
	genesisBlock := dimos.Block{
		IDx:          1,
		Timestamp:    time.Now().Unix(),
		Transactions: []dimos.Transaction{},
		PrevHash:     []byte("0"),
		Signature:    []byte("0"),
	}

	// Get the merkle root.
	root, err := genesisBlock.ComputeMerkleRoot()

	log.Println("Merkle Root: ", hex.EncodeToString(root), err)

	// Compute the hash of the block.
	hash, err := genesisBlock.ComputeHash()

	log.Println("Genesis Hash:", hex.EncodeToString(hash), err)

	// Write the genesis hash to the disk.
	_ = utils.WriteGenesisHash(hash)

	// Add the block.
	blockchain.AddBlock(&genesisBlock)

	log.Println("Created genesis")
}
