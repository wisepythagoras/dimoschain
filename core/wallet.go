package core

import (
	"github.com/vmihailenco/msgpack/v4"
	"github.com/wisepythagoras/dimoschain/crypto"
)

// Wallet defines the wallet structure.
type Wallet struct {
	KeyPair *crypto.KeyPair
}

// Serialize serializes the wallet
func (w *Wallet) Serialize() []byte {
	return w.KeyPair.Private.Serialize()
}

// Unserialize imports a wallet file.
func (w *Wallet) Unserialize(bin []byte) error {
	privKeyBytes := make([]byte, len(bin))

	// Unmarshal the binary into a private key.
	err := msgpack.Unmarshal(bin, &privKeyBytes)

	if err != nil {
		return err
	}

	// Parse the private key from bytes.
	w.KeyPair = crypto.PrivKeyFromBytes(privKeyBytes)

	return nil
}
