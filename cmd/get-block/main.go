package main

import (
	"encoding/hex"
	"flag"
	"log"

	"github.com/wisepythagoras/dimoschain/dimos"
)

func main() {
	// Define and parse the command line arguments.
	hash := flag.String("hash", "", "The hash of the block")
	getCurrent := flag.Bool("current", false, "Get the current block")

	flag.Parse()

	if !*getCurrent && (hash == nil || *hash == "") {
		log.Fatal("Either -hash <hash> or -current is required")
		return
	}

	// Load the database.
	blockchain, err := dimos.InitChainDB()

	if err != nil {
		log.Fatal(err)
		return
	}

	var block *dimos.Block

	if *getCurrent {
		block, err = blockchain.GetCurrentBlock()
	} else {
		decodedHash, _ := hex.DecodeString(*hash)
		block, err = blockchain.GetBlock(decodedHash)
	}

	log.Println(block, err)
}
