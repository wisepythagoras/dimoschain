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

	log.Println("Local chain current block", hex.EncodeToString(blockchain.CurrentHash))

	// Start the server.
	server := &Server{
		Port:       8013,
		Blockchain: blockchain,
	}

	// Start listening.
	server.Listen()
}
