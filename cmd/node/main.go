package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"net"

	"github.com/cossacklabs/themis/gothemis/keys"
	"github.com/cossacklabs/themis/gothemis/session"
	"github.com/wisepythagoras/dimoschain/core"
	"github.com/wisepythagoras/dimoschain/utils"
)

type callbacks struct{}

func (clb *callbacks) GetPublicKeyForId(ss *session.SecureSession, id []byte) *keys.PublicKey {
	decodedID, err := base64.StdEncoding.DecodeString(string(id[:]))
	if nil != err {
		return nil
	}
	return &keys.PublicKey{Value: decodedID}
}

func (clb *callbacks) StateChanged(ss *session.SecureSession, state int) {

}

func clientHandler(c net.Conn, serverID string, serverPrivateKey *keys.PrivateKey) {
	// Create a secure session.
	secureSession, err := session.New([]byte(serverID), serverPrivateKey, &callbacks{})

	if err != nil {
		log.Fatalln(err)
	}

	for {
		buf := make([]byte, 8192)

		// Get the bytes from the received message and write them to our buffer.
		readBytes, err := c.Read(buf)

		if err != nil {
			log.Fatalln(err)
		}

		buf, sendPeer, err := secureSession.Unwrap(buf[:readBytes])

		if nil != err {
			log.Fatalln(err)
		}

		if !sendPeer {
			if "finish" == string(buf[:]) {
				return
			}

			fmt.Println("Received:", string(buf[:]))
			buf, err = secureSession.Wrap(buf)

			if nil != err {
				log.Fatalln(err)
			}
		}

		_, err = c.Write(buf)

		if err != nil {
			log.Fatalln(err)
		}
	}
}

func server(blockchain *core.Blockchain) {
	// Start a TCP server that listens on port 8013.
	server, err := net.Listen("tcp", ":8013")

	if err != nil {
		log.Fatalln(err)
	}

	// Here we generate a new keypair for the server to use. We may want to use existing EC keys.
	serverKeyPair, err := keys.New(keys.TypeEC)

	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := server.Accept()

		if err != nil {
			log.Fatalln(err)
		}

		// Handle each client in its own thread.
		go clientHandler(conn, base64.StdEncoding.EncodeToString(serverKeyPair.Public.Value), serverKeyPair.Private)
	}
}

func main() {
	log.Println(utils.Name, utils.Version)

	// Load the blockchain database.
	blockchain, err := core.InitChainDB()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Current state block hash is", hex.EncodeToString(blockchain.CurrentHash))

	// Start the server.
	server(blockchain)
}
