package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/cossacklabs/themis/gothemis/session"
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

	rw, err := c.createInstanceAndConnect(ctx, h, c.Address)

	if err != nil {
		return err
	}

	// Create a thread to read and write data.
	go writeData(rw)
	go readData(rw)

	// Wait forever
	select {}
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
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	return rw, nil
}

// ConnectToNode connects to the node with the specified IP address and port.
/*func ConnectToNode(ip string, port int) (*Connection, error) {
	// Validate the cnputs.
	if !utils.IsIPAddressValid(ip) || port <= 0 || port > 65535 {
		return nil, errors.New("Invalid IP address or port")
	}

	// The first thing we want to do is connect to the TCP port of the node.
	conn, err := net.Dial("tcp", ip+":"+strconv.FormatInt(int64(port), 10))

	if err != nil {
		return nil, err
	}

	// Generate the keys necessary to communicate with the node.
	clientKeyPair, err := keys.New(keys.TypeEC)

	if err != nil {
		return nil, err
	}

	pubKey := base64.StdEncoding.EncodeToString(clientKeyPair.Public.Value)

	// This is what initiates a secure session.
	secureSession, err := session.New([]byte(pubKey), clientKeyPair.Private, &proto.Callback{})

	if err != nil {
		return nil, err
	}

	buf, err := secureSession.ConnectRequest()

	// The following for loop is taken from the `secure_session_client.go` example.
	// TODO: Why is this necessary?
	for {
		_, err = conn.Write(buf)

		if err != nil {
			return nil, err
		}

		buf = make([]byte, 10240)
		readBytes, err := conn.Read(buf)

		if err != nil {
			return nil, err
		}

		buffer, sendPeer, err := secureSession.Unwrap(buf[:readBytes])

		if nil != err {
			return nil, err
		}

		buf = buffer

		if !sendPeer {
			break
		}
	}

	if err != nil {
		return nil, err
	}

	if err != nil {
		fmt.Println("raw message error")
		return nil, err
	}

	connection := &Connection{
		secureSession: secureSession,
		connection:    conn,
		ip:            ip,
		port:          port,
	}

	return connection, nil
}*/

// Connection is the type that handles a connection with a node.
type Connection struct {
	ip            string
	port          int
	secureSession *session.SecureSession
	connection    net.Conn
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

	// This encrypts the message that's sent to the node.
	buf, err := c.secureSession.Wrap(packed)

	if err != nil {
		return false
	}

	// Now we send the encrypted message over to the node.
	_, err = c.connection.Write(buf)

	if err != nil {
		return false
	}

	return true
}
