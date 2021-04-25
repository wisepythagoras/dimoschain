package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/cossacklabs/themis/gothemis/keys"
	"github.com/cossacklabs/themis/gothemis/session"
	"github.com/wisepythagoras/dimoschain/proto"
	"github.com/wisepythagoras/dimoschain/utils"
)

// ConnectToNode connects to the node with the specified IP address and port.
func ConnectToNode(ip string, port int) (*Connection, error) {
	// Validate the inputs.
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
}

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
