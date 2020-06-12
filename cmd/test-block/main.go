package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/wisepythagoras/dimoschain/core"
	"github.com/wisepythagoras/dimoschain/crypto"
)

func genKeyPair() *crypto.KeyPair {
	// Create a new key pair instance.
	keyPair := crypto.KeyPair{}

	// Generate the new keypair.
	err := keyPair.Generate()

	if err != nil {
		log.Fatalln(err)
	}

	return &keyPair
}

func main() {
	// Load the database.
	blockchain, err := core.InitChainDB()

	if err != nil {
		log.Fatal(err)
		return
	}

	// First we need to get the current block. This will determine which IDx this
	// new block will take, as well as the previous hash.
	currentBlock, err := blockchain.GetCurrentBlock()

	if err != nil {
		log.Fatal(err)
		return
	}

	sender := genKeyPair()
	receiver := genKeyPair()

	fmt.Println(sender.GetAddr(), "->", receiver.GetAddr())

	// Create a test transaction.
	tx := core.Transaction{
		Hash:      nil,
		Amount:    0.001,
		From:      []byte(sender.GetAddr()),
		To:        []byte(receiver.GetAddr()),
		Signature: []byte("test1signature"),
	}

	// Construct the new block.
	block := core.Block{
		IDx:          currentBlock.IDx + 1,
		Hash:         nil,
		PrevHash:     currentBlock.Hash,
		MerkleRoot:   nil,
		Timestamp:    time.Now().Unix(),
		Transactions: []core.Transaction{},
		Signature:    []byte("test"),
	}

	// Adding a transaction will also change the hash on the block.
	block.AddTransaction(&tx)

	// But we run this anyway, since a block could be empty.
	block.UpdateHash()

	// Finally add the block.
	success, err := blockchain.AddBlock(&block)

	if err != nil {
		log.Fatal(err)
	} else if !success {
		log.Fatal("Unale to add the new block")
	} else {
		hexHash := hex.EncodeToString(block.Hash)
		msg := fmt.Sprintf("The block was added with hash %s", hexHash)
		log.Println(msg)
	}
}
