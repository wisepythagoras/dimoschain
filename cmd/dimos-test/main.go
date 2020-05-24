package main

import (
	"fmt"
	"encoding/hex"
	"github.com/wisepythagoras/dimoschain/dimos"
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

	fmt.Println(crypto.GenAddr(keyPair.Public.SerializeCompressed()))

	tx := dimos.Transaction{}

	fmt.Println(tx)
}
