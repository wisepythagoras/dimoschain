package proto

import (
	"github.com/vmihailenco/msgpack"
)

// Message defines what the protocol's message structure.
type Message struct {
	Command int `json:"command" msgpack:"c"`

	// In order to make this struct flexible and expand its use, we allow the payload
	// to be any arbitrary type.
	Payload interface{} `json:"payload" msgpack:"p"`
}

// Pack creates the msgpack version of the message.
func (m *Message) Pack() ([]byte, error) {
	return msgpack.Marshal(m)
}

// Unpack unmarshals the msgpack encoded message into an instance of a Message.
func Unpack(bytes []byte) (*Message, error) {
	// This is the instance which will contain the results of the unpacking.
	msg := &Message{}

	// Try decoding the given message.
	err := msgpack.Unmarshal(bytes, msg)

	if err != nil {
		return nil, err
	}

	return msg, nil
}
