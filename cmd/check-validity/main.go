package main

import (
	"log"

	"github.com/wisepythagoras/dimoschain/dimos"
)

func main() {
	// Load the database.
	blockchain, err := dimos.InitChainDB()

	if err != nil {
		log.Fatal(err)
		return
	}

	// Validate the blockchain.
	success, err := blockchain.IsChainValid(true)

	log.Println(success, err)
}
