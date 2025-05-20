package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/multiformats/go-multiaddr"
	"github.com/wisepythagoras/dimoschain/core"
	"github.com/wisepythagoras/dimoschain/proto"
)

// Server defines the server struct.
type Client struct {
	Port       int
	Address    string
	Blockchain *core.Blockchain
	randomness io.Reader
	readWriter *bufio.ReadWriter
}

func (c *Client) Create() error {
	ctx, cancel := context.WithCancel(context.Background())
	_ = ctx
	defer cancel()

	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.Ed25519, 256, c.randomness)

	if err != nil {
		return err
	}

	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", c.Port))

	h, err := libp2p.New(
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
	)

	if err != nil {
		return err
	}

	_, err = c.createInstanceAndConnect(ctx, h, c.Address)

	if err != nil {
		return err
	}

	// Create a thread to read and write data.
	go writeData(c.readWriter)
	go readData(c.readWriter)

	// Wait forever
	select {}
}

func (c *Client) SendData(data []byte) error {
	_, err := c.readWriter.Write(data)

	if err != nil {
		return err
	}

	return c.readWriter.Flush()
}

func (c *Client) createInstanceAndConnect(ctx context.Context, h host.Host, destination string) (*bufio.ReadWriter, error) {
	log.Println("Node multiaddress:")

	for _, la := range h.Addrs() {
		log.Printf(" - %v\n", la)
	}

	log.Println()

	maddr, err := multiaddr.NewMultiaddr(destination)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	info, err := peer.AddrInfoFromP2pAddr(maddr)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	h.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

	s, err := h.NewStream(context.Background(), info.ID, "/dimos/1.0.0")

	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create a buffered stream so that read and writes are non-blocking.
	c.readWriter = bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	return c.readWriter, nil
}

// Connection is the type that handles a connection with a node.
type Connection struct {
	ip     string
	port   int
	Client *Client
}

// SendCommand sends a command to
func (c *Connection) SendCommand(command int, payload interface{}) bool {
	msg := proto.Message{
		Command: command,
		Payload: payload,
	}

	packed, err := msg.Pack()

	if err != nil {
		return false
	}

	err = c.Client.SendData(packed)

	if err != nil {
		return false
	}

	return true
}
