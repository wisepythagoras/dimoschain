package main

import (
	"log"

	"github.com/wisepythagoras/dimoschain/core"
)

func main() {
	// Load the database.
	blockchain, err := core.InitChainDB()

	if err != nil {
		log.Fatal(err)
		return
	}

	// Validate the blockchain.
	success, err := blockchain.IsChainValid(true)

	log.Println(success, err)
}
