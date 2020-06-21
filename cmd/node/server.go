package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/cossacklabs/themis/gothemis/keys"
	"github.com/cossacklabs/themis/gothemis/session"
)

// Server defines the server struct.
type Server struct {
	Port int
}

// Listen starts the server and listens on the designated port.
func (s *Server) Listen() {
	// Start a TCP server that listens on port 8013.
	server, err := net.Listen("tcp", ":"+strconv.Itoa(s.Port))

	if err != nil {
		log.Fatalln(err)
	}

	// Here we generate a new keypair for the server to use. We may want to use existing EC keys.
	serverKeyPair, err := keys.New(keys.TypeEC)

	if err != nil {
		log.Fatalln(err)
	}

	for {
		// Here we accept a new connection.
		conn, err := server.Accept()

		if err != nil {
			log.Fatalln(err)
		}

		// Encode the public key.
		pubKey := base64.StdEncoding.EncodeToString(serverKeyPair.Public.Value)

		// Handle each client in its own thread.
		go s.clientHandler(conn, pubKey, serverKeyPair.Private)
	}
}

// clientHandler handles any incoming connection. It's not exported on purpose, since its use is only
// internal.
func (s *Server) clientHandler(c net.Conn, serverID string, serverPrivateKey *keys.PrivateKey) {
	// Create a secure session.
	secureSession, err := session.New([]byte(serverID), serverPrivateKey, &Callback{})

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