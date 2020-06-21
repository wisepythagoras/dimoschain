package proto

import (
	"testing"
)

// TestPackUnpack tests whether packing and unpacking works as expected.
func TestPackUnpack(t *testing.T) {
	// Create the test message.
	msg := Message{
		Command: "test",
		Payload: 123,
	}

	// Pack the message.
	packed, err := msg.Pack()

	// Ensure there were no errors.
	if err != nil {
		t.Errorf("Unexpected error during pack: %s", err)
	}

	// Unpack the message.
	unpacked, err := Unpack(packed)

	// Ensure there were no errors.
	if err != nil {
		t.Errorf("Unexpected error during unpack: %s", err)
	}

	// Now assert that the right data has been decoded.
	if p, _ := msg.Payload.(int); unpacked.Command != msg.Command || unpacked.Payload != int64(p) {
		t.Log("unpacked.Command != msg.Command", unpacked.Command != msg.Command)
		t.Log("unpacked.Payload != msg.Payload", unpacked.Payload != msg.Payload)
		t.Errorf("Invalid unpack \"%s\", \"%s\"/\"%s\", \"%s\"", unpacked.Command, unpacked.Payload, msg.Command, msg.Payload)
	}
}
