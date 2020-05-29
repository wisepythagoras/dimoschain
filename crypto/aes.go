package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

var IV = []byte{212, 126, 197, 12, 20, 238, 61, 80, 12, 162, 12, 45, 227, 27, 150, 43}

// PadKey compensates for a key that is smaller than 32 bytes.
func PadKey(key []byte) []byte {
	// Nothing to do here.
	if len(key) == 32 {
		return key
	}

	// We will fill the key with random bytes.
	prbg := PRBG{
		Seed: key,
	}

	// Recompose the key.
	return append(key, prbg.Next(32-len(key))...)
}

// EncryptCTR encrypts the plaintext with AES/CTR.
func EncryptCTR(plaintext []byte, key []byte) ([]byte, error) {
	// Create a new cipher.
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	// Now we make the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// The IV needs to be unique and random, but, according to the docs, it doesn't need
	// to be secure.
	iv := ciphertext[:aes.BlockSize]

	// So let's get some random values from the reader.
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Now let's encrypt.
	stream := cipher.NewCTR(block, IV)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// Should an HMAC be used here?

	// Return the ciphertext.
	return ciphertext, nil
}

// DecryptCTR decrypts an AES/CTR encrypted ciphertext.
func DecryptCTR(ciphertext []byte, key []byte) ([]byte, error) {
	// Create a new cipher.
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	iv := ciphertext[:aes.BlockSize]

	// So let's get some random values from the reader.
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	plaintext := make([]byte, aes.BlockSize+len(ciphertext))

	// Now let's decrypt.
	stream := cipher.NewCTR(block, IV)
	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])

	return plaintext, nil
}
