package crypto

import (
	"testing"
)

// TestSHA256 tests that the SHA-256 function returns the correct data.
func TestSHA256(t *testing.T) {
	// Generate the SHA-256 sum
	hash := GetSHA256Hash([]byte("test"))
	hexHash := ByteArrayToHex(hash)
	targetHash := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"

	if hexHash != targetHash {
		t.Error("Invalid result hash")
	}
}

// TestSHA3384 tests that the SHA-384 function returns the correct data.
func TestSHA3384(t *testing.T) {
	// Generate the SHA-384 sum
	hash, err := GetSHA3384Hash([]byte("test"))

	if err != nil {
		t.Error(err)
	}

	hexHash := ByteArrayToHex(hash)
	targetHash := "e516dabb23b6e30026863543282780a3ae0dccf05551cf0295178d7ff0f1b41eecb9db3ff219007c4e097260d58621bd"

	if hexHash != targetHash {
		t.Error("Invalid result hash")
	}
}

// TestBase58Encoding encodes a string a string in base58.
func TestBase58Encoding(t *testing.T) {
	str := "test"

	// Encode the string.
	enc := Base58Encode([]byte(str))

	if enc != "3yZe7d" {
		t.Error("The string was not base58 encoded correctly")
	}

	// Decode the string.
	dec := Base58Decode(enc)

	if string(dec) != str {
		t.Error("The string was not base58 decoded correctly")
	}
}
