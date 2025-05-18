package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/multiformats/go-multiaddr"
	"github.com/wisepythagoras/dimoschain/core"
)

// Server defines the server struct.
type Server struct {
	Port       int
	Blockchain *core.Blockchain
	randomness io.Reader
}

func (s *Server) Create() error {
	ctx, cancel := context.WithCancel(context.Background())
	_ = ctx
	defer cancel()

	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.Ed25519, 256, s.randomness)

	if err != nil {
		return err
	}

	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", s.Port))

	h, err := libp2p.New(
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
	)

	if err != nil {
		return err
	}

	h.SetStreamHandler("/dimos/1.0.0", s.handleStream)

	log.Printf("Connect to: \"/ip4/127.0.0.1/tcp/%v/p2p/%s\"\n", s.Port, h.ID())
	log.Println("Instead of 127.0.0.1, use a public IP as well")
	log.Println()

	// Wait forever
	select {}
}

func (s *Server) handleStream(st network.Stream) {
	log.Println("Got a new stream!")

	// Create a buffer stream for non-blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(st), bufio.NewWriter(st))

	go readData(rw)
	go writeData(rw)

	// st.Close()
}

/*
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
			fmt.Println(err)
			continue
		}

		// Encode the public key.
		pubKey := base64.StdEncoding.EncodeToString(serverKeyPair.Public.Value)

		// Handle each client in its own thread.
		// go s.clientHandler(conn, pubKey, serverKeyPair.Private)
		_, _ = conn, pubKey
	}
}


// clientHandler handles any incoming connection. It's not exported on purpose, since its use is only
// internal.
func (s *Server) clientHandler(c net.Conn, serverID string, serverPrivateKey *keys.PrivateKey) {
	// Create a secure session.
	secureSession, err := session.New([]byte(serverID), serverPrivateKey, &proto.Callback{})

	if err != nil {
		log.Println(err)
		return
	}

	for {
		buf := make([]byte, 8192)

		// Get the bytes from the received message and write them to our buffer.
		readBytes, err := c.Read(buf)

		if err != nil {
			log.Println("Net error", err)
			return
		}

		// Decrypt the encrypted data from the peer.
		buf, sendPeer, err := secureSession.Unwrap(buf[:readBytes])

		if err != nil {
			log.Println(err)
			continue
		}

		if !sendPeer {
			// Unpack the message.
			message, err := proto.Unpack(buf)

			if err != nil {
				log.Println(err)
				continue
			}

			// Was the exit command received?
			if message.Command == proto.CmdExit {
				return
			}

			log.Printf("Cmd: %d: %s\n", message.Command, message.Payload)

			// Echo for now.
			buf, err = secureSession.Wrap(buf)

			if err != nil {
				log.Println(err)
				continue
			}
		}

		_, err = c.Write(buf)

		if err != nil {
			log.Println("End", err)
			continue
		}
	}
}*/
