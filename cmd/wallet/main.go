package main

import (
	"flag"
	"log"

	"github.com/vmihailenco/msgpack"
	"github.com/wisepythagoras/dimoschain/crypto"
	"github.com/wisepythagoras/dimoschain/dimos"
	"github.com/wisepythagoras/dimoschain/utils"
)

func main() {
	// Define and parse the command line arguments.
	privKey := flag.String("import", "", "Import a private key")
	open := flag.String("open", "", "Open a wallet file")
	create := flag.Bool("create", false, "Create a new wallet")
	name := flag.String("name", "", "The filename")

	flag.Parse()

	if *privKey == "" && *open == "" && *create == false {
		log.Println("A private key is needed (-import) or open a wallet file (-open)")
		log.Fatalln("Otherwise create a new wallet (-create)")
	}

	if *create {
		// Create a new key pair instance.
		keyPair := crypto.KeyPair{}

		// Generate the new keypair.
		err := keyPair.Generate()

		if err != nil {
			log.Fatalln(err)
			return
		}

		// Get the address.
		addr := keyPair.GetAddr()

		log.Println("New Address:", addr)
		log.Println("Public Key: ", keyPair.GetPubKey())
		log.Println("Private Key:", keyPair.GetPrivKey())

		if len(*name) > 0 {
			// Create a new keypair object.
			wallet := dimos.Wallet{
				KeyPair: keyPair,
			}

			// Marshall to msgpack.
			bin, _ := msgpack.Marshal(wallet.Serialize())

			// Write the file to the disk.
			err = utils.WriteToFile(*name+".wallet", bin)
		}

		if err != nil {
			log.Fatalln(err)
		}
	} else {
		log.Println("Under construction")
	}
}
