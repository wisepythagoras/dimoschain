package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"syscall"

	"github.com/vmihailenco/msgpack"
	"github.com/wisepythagoras/dimoschain/crypto"
	"github.com/wisepythagoras/dimoschain/dimos"
	"github.com/wisepythagoras/dimoschain/utils"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	// Define and parse the command line arguments.
	privKey := flag.String("import", "", "Import a private key")
	open := flag.String("open", "", "Open a wallet file")
	create := flag.Bool("create", false, "Create a new wallet")
	name := flag.String("name", "", "The filename")
	protect := flag.Bool("protect", false, "Whether to protect the wallet with a password")

	flag.Parse()

	if *privKey == "" && *open == "" && *create == false {
		log.Println("A private key is needed (-import) or open a wallet file (-open)")
		log.Println("Otherwise create a new wallet (-create)")
		log.Println("If you want to save the new keys to a wallet file, then pass the -name <name> flag")
		log.Fatalln("You can also password encrypt the wallet with -protect")
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

		fmt.Println("New Address:", addr)
		fmt.Println("Public Key: ", keyPair.GetPubKey())
		fmt.Println("Private Key:", keyPair.GetPrivKey())

		if len(*name) > 0 {
			// Create a new keypair object.
			wallet := dimos.Wallet{
				KeyPair: keyPair,
			}

			// Marshall to msgpack.
			bin, _ := msgpack.Marshal(wallet.Serialize())

			if *protect {
				var password []byte
				gotPassword := false

				for !gotPassword {
					fmt.Print("Enter Password: ")

					// Read the password
					tempPassword, err := terminal.ReadPassword(int(syscall.Stdin))

					if err != nil {
						log.Fatalln(err)
					}

					fmt.Println("")
					fmt.Print("Confirm Password: ")

					// Now read the confirmation.
					confirmPassword, err := terminal.ReadPassword(int(syscall.Stdin))

					fmt.Println("")

					if err != nil {
						log.Fatalln(err)
					}

					if bytes.Compare(tempPassword, confirmPassword) != 0 {
						fmt.Println("The passwords didn't match. Try again.")
					} else {
						gotPassword = true
						password = tempPassword
					}
				}

				// Encrypt the payload with the password.
				if bin, err = crypto.EncryptGCM(bin, crypto.PadKey(password)); err != nil {
					log.Fatal(err)
				}
			}

			// Write the file to the disk.
			err = utils.WriteToFile(*name+".wallet", bin)
		}

		if err != nil {
			log.Fatalln(err)
		}
	} else {
		fmt.Println("Under construction")
	}
}
