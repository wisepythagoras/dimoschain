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

	// TODO: This logic will be combined. It's here just for testing purposes.
	go readData(rw)
	go writeData(rw)

	// st.Close()
}
