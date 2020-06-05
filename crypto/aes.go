package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

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

// EncryptGCM encrypts the plaintext with AES/GCM.
func EncryptGCM(plaintext []byte, key []byte) ([]byte, error) {
	// Create a new cipher.
	c, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	// Use GCM - Galois/Counter Mode.
	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return nil, err
	}

	// GCM requires that we have a cryptographically secure nonce which will be passed
	// to our Seal function.
	nonce := make([]byte, gcm.NonceSize())

	// Get the nonce.
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Now seal the deal.
	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// DecryptGCM decrypts an AES/GCM encrypted ciphertext.
func DecryptGCM(ciphertext []byte, key []byte) ([]byte, error) {
	// Create a new cipher.
	c, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	// Use GCM - Galois/Counter Mode.
	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	if len(ciphertext) < nonceSize {
		return nil, errors.New("The size of the nonce is greater than the length of the ciphertext")
	}

	// Separate the nonce from the ciphertext.
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Now decrypt the ciphertext.
	return gcm.Open(nil, nonce, ciphertext, nil)
}
