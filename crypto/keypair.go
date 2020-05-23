package crypto

import (
	"errors"
	"encoding/hex"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/decred/dcrd/dcrec/secp256k1/ecdsa"
)

// KeyPair represents a structure for managing encryption keys.
type KeyPair struct {
	Public *secp256k1.PublicKey
	Private *secp256k1.PrivateKey
}

// Generate generates a new set of secp256k1 keys.
func (k *KeyPair) Generate() error {
	// Generate a private key.
	key, err := secp256k1.GeneratePrivateKey()

	if err != nil {
		return err
	}

	// Set the keys on the local object.
	k.Private = key
	k.Public = key.PubKey()

	return nil
}

// Sign simply signs a message.
func (k *KeyPair) Sign(message []byte) (*ecdsa.Signature, error) {
	if k.Private == nil {
		return nil, errors.New("No private key loaded")
	}

	// Get the SHA3-512 hash of the message.
	hash, err := GetSHA3512Hash(message)

	if err != nil {
		return nil, err
	}

	// Sign the message.
	return ecdsa.Sign(k.Private, hash), nil
}

// GetPubKey returns the public key.
func (k *KeyPair) GetPubKey() string {
	return hex.EncodeToString(k.Public.SerializeCompressed())
}

// GetAddr gets the address version of the public key.
func (k *KeyPair) GetAddr() string {
	// Get the public key bytes.
	public := k.Public.SerializeCompressed()

	var addr []byte

	// Add the version to the address.
	addr = append([]byte{0x01}, public...)

	// Return the wallet address.
	return Base58Encode(addr)
}

func (k *KeyPair) GetPubKeyFromAddr(str string) error {
	// Decode the base58 encoded string.
	decoded := Base58Decode(str)

	if len(decoded) == 0 || len(decoded) < 34 {
		return errors.New("Invalid input string")
	}

	return nil
}

// GetPubKeyUncompressed gets the uncompressed version of the public key.
func (k *KeyPair) GetPubKeyUncompressed() string {
	return hex.EncodeToString(k.Public.SerializeUncompressed())
}

// GetPrivKey returns the private key.
func (k *KeyPair) GetPrivKey() string {
	return hex.EncodeToString(k.Private.Serialize())
}

// PrivKeyFromBytes gets the private key from bytes.
func PrivKeyFromBytes(priv []byte) *KeyPair {
	// Parse the private key.
	private := secp256k1.PrivKeyFromBytes(priv)

	// Return the new key pair.
	return &KeyPair{
		private.PubKey(),
		private,
	}
}

// ParsePubKey parses a pubic key.
func ParsePubKey(pub []byte) (*KeyPair, error) {
	pubkey, err := secp256k1.ParsePubKey(pub)

	if err != nil {
		return nil, err
	}

	return &KeyPair{pubkey, nil}, nil
}

// VerifySignature verifies a DER signature.
func VerifySignature(pub *secp256k1.PublicKey, sig []byte, msg []byte) bool {
	// Parse the DER signature.
	signature, err := ecdsa.ParseDERSignature(sig)

	if err != nil {
		return false
	}

	// Get the SHA3-512 hash of the message.
	hash, err := GetSHA3512Hash(msg)

	if err != nil {
		return false
	}

	// Verify the signature.
	return signature.Verify(hash, pub)
}
