package main

import (
	"log"

	"github.com/wisepythagoras/dimoschain/core"
	"github.com/wisepythagoras/dimoschain/utils"
)

func main() {
	log.Println(utils.Name, utils.Version)

	// Load the database.
	blockchain, err := core.InitChainDB()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(blockchain)
}
