package main

import (
	"encoding/hex"
	"log"

	"github.com/wisepythagoras/dimoschain/core"
	"github.com/wisepythagoras/dimoschain/utils"
)

func main() {
	log.Println(utils.Name, utils.Version)

	// Load the blockchain database.
	blockchain, err := core.InitChainDB()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Current state block hash is", hex.EncodeToString(blockchain.CurrentHash))

	// Start the server.
	server := &Server{
		Port: 8013,
	}

	// Start listening.
	server.Listen()
}
