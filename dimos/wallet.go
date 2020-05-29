package dimos

import (
	"github.com/vmihailenco/msgpack"
	"github.com/wisepythagoras/dimoschain/crypto"
)

// Wallet defines the wallet structure.
type Wallet struct {
	KeyPair crypto.KeyPair
}

// Serialize serializes the wallet
func (w *Wallet) Serialize() []byte {
	wallet := make(map[string]interface{})
	wallet["pub"] = w.KeyPair.Public.SerializeCompressed()
	wallet["priv"] = w.KeyPair.Private.Serialize()

	// Serialize the object into msgpack format.
	serialized, _ := msgpack.Marshal(wallet)

	return serialized
}
