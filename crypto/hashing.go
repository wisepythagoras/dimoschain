package crypto

import (
	"encoding/hex"
	"hash"

	"github.com/btcsuite/btcd/btcutil/base58"
	"golang.org/x/crypto/sha3"

	"crypto/sha256"
)

var (
	// HashStrategy is the hash strategy for SHA3-384.
	HashStrategy func() hash.Hash = sha3.New384
)

// Base58Encode encodes a series of bytes into a base58 string.
func Base58Encode(payload []byte) string {
	return base58.Encode(payload)
}

// Base58Decode decodes a base58 encoded string.
func Base58Decode(str string) []byte {
	return base58.Decode(str)
}

// GetSHA3384Hash returns the SHA3-512 hash of a given string.
func GetSHA3384Hash(str []byte) ([]byte, error) {
	// Create a new sha object.
	h := HashStrategy()

	// Add our string to the hash.
	if _, err := h.Write([]byte(str)); err != nil {
		return nil, err
	}

	// Return the SHA3-512 digest.
	return h.Sum(nil), nil
}

// GetSHA256Hash returns the SHA256 hash.
func GetSHA256Hash(b []byte) []byte {
	sha256 := sha256.New()
	sha256.Write(b)
	return sha256.Sum(nil)
}

// ByteArrayToHex converts a set of bytes to a hex encoded string.
func ByteArrayToHex(payload []byte) string {
	return hex.EncodeToString(payload)
}

// DoubleSHA256 generates the double SHA256 hash of the input.
func DoubleSHA256(b []byte) []byte {
	return GetSHA256Hash(GetSHA256Hash(b))
}

// AddrFromPubKey generates an address the same way it's generated for Bitcoin.
func AddrFromPubKey(pubkey []byte) string {
	a := append([]byte{0}, pubkey...)

	return Base58Encode(append(a, DoubleSHA256(a)...))
}
