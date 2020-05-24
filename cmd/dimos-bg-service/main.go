package main

import (
	"log"

	"github.com/wisepythagoras/dimoschain/dimos"
	"github.com/wisepythagoras/dimoschain/utils"
)

func main() {
	log.Println(utils.NAME, utils.VERSION)

	// Load the database.
	blockchain, err := dimos.InitChainDB()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(blockchain)
}
