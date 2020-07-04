package main

import (
	"encoding/hex"
	"fmt"

	"github.com/wisepythagoras/dimoschain/core"
	"github.com/wisepythagoras/dimoschain/crypto"
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
	keyPair.GetPubKeyHashFromAddr(addr)

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

	tx := core.Transaction{}

	fmt.Println(tx)

	prbg := crypto.PRBG{
		Seed: []byte("Test seed"),
	}

	fmt.Println(hex.EncodeToString(prbg.Next(10)))
	fmt.Println(hex.EncodeToString(prbg.Next(10)))
	fmt.Println(prbg.NextUInt64(10))

	fmt.Println("----")

	plaintext := []byte("This is an example plaintext")
	key := []byte("This is a test key")

	ciphertext, err := crypto.EncryptGCM(plaintext, crypto.PadKey(key))

	fmt.Println("Ciphertext:", hex.EncodeToString(ciphertext), err)

	decrypted, err := crypto.DecryptGCM(ciphertext, crypto.PadKey(key))

	fmt.Println("Decrypted:", string(decrypted), err)
}
