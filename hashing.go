package main

import (
	"encoding/hex"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/sha3"

	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)

// Base58Encde encodes a series of bytes into a base58 string.
func Base58Encode(payload []byte) string {
	return base58.Encode(payload)
}

// Base58Decode decodes a base58 encoded string.
func Base58Decode(str string) []byte {
        return base58.Decode(str)
}

// GetSHA3512Hash returns the SHA3-512 hash of a given string.
func GetSHA3512Hash(str string) ([]byte, error) {
	// Create a new sha object.
	h := sha3.New512()

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

// Ripemd160SHA256 generates a Ripemd160 hash which is used for wallet addresses.
func Ripemd160SHA256(b []byte) []byte {
	shaSum := GetSHA256Hash(b)

	hash := ripemd160.New()
	hash.Write(shaSum)

	return hash.Sum(nil)
}

// GenArr generates an address the same way it's generated for Bitcoin.
func GenAddr(pubkey []byte) string {
	ripe := Ripemd160SHA256(pubkey)
	a := append([]byte{0}, ripe...)

	return Base58Encode(append(a, GetSHA256Hash(GetSHA256Hash(a))[:4]...))
}
