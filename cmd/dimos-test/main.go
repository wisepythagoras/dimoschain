package main

import (
	"encoding/hex"
	"fmt"

	"github.com/wisepythagoras/dimoschain/crypto"
	"github.com/wisepythagoras/dimoschain/dimos"
)

func main() {
	fmt.Println("Hello, world!")
	//GenerateKey()

	keyPair := crypto.KeyPair{}
	err := keyPair.Generate()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Public:", keyPair.GetPubKey())
	fmt.Println("Private:", keyPair.GetPrivKey())

	addr := keyPair.GetAddr()

	fmt.Println("Address:", addr)
	keyPair.GetPubKeyFromAddr(addr)

	sig, err := keyPair.Sign([]byte("Hello, world!"))

	if err != nil {
		fmt.Println(err)
		return
	}

	der := sig.Serialize()
	fmt.Println("Signature: ", hex.EncodeToString(der))
	res := crypto.VerifySignature(keyPair.Public, der, []byte("Hello, world!"))
	fmt.Println("Verified:", res)

	fmt.Println(crypto.AddrFromPubKey(keyPair.Public.SerializeCompressed()))

	tx := dimos.Transaction{}

	fmt.Println(tx)

	prbg := crypto.PRBG{
		Seed: []byte("Test seed"),
	}

	fmt.Println(hex.EncodeToString(prbg.Next(10)))
	fmt.Println(hex.EncodeToString(prbg.Next(10)))
	fmt.Println(prbg.NextInt64(10))

	fmt.Println("----")

	plaintext := []byte("This is an example plaintext")
	key := []byte("This is a test key")

	ciphertext, err := crypto.EncryptCTR(plaintext, crypto.PadKey(key))

	fmt.Println("Ciphertext:", hex.EncodeToString(ciphertext), err)

	decrypted, err := crypto.DecryptCTR(ciphertext, crypto.PadKey(key))

	fmt.Println("Decrypted:", string(decrypted), err)
}
