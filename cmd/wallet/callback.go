package main

import (
	"encoding/base64"

	"github.com/cossacklabs/themis/gothemis/keys"
	"github.com/cossacklabs/themis/gothemis/session"
)

// Callback defines a custom callbacks type.
type Callback struct{}

// GetPublicKeyForId returns the public key for the id.
func (clb *Callback) GetPublicKeyForId(ss *session.SecureSession, id []byte) *keys.PublicKey {
	// Decode the id.
	decodedID, err := base64.StdEncoding.DecodeString(string(id[:]))

	if nil != err {
		return nil
	}

	// Return the public key for the decoded id.
	return &keys.PublicKey{Value: decodedID}
}

// StateChanged handles a state change.
func (clb *Callback) StateChanged(ss *session.SecureSession, state int) {
	// Pass through.
}
