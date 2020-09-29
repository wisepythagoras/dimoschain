GOPATH := $(shell go env GOPATH)
SRC := $(GOPATH)/src/github.com/wisepythagoras/

all: validity wallet get-block genesis bg-service node

wallet:
	make -C cmd/wallet

bg-service:
	make -C cmd/background-service

genesis:
	make -C cmd/create-genesis

get-block:
	make -C cmd/get-block

validity:
	make -C cmd/check-validity

node:
	make -C cmd/node

tests:
	make -C cmd/test-block
	make -C cmd/dimos-test

install-deps:
	go get github.com/cbergoon/merkletree
	go get github.com/decred/dcrd/dcrec/secp256k1
	go get github.com/btcsuite/btcutil/base58
	go get golang.org/x/crypto/sha3
	go get golang.org/x/crypto/ripemd160
	go get github.com/dgraph-io/badger
	go get github.com/vmihailenco/msgpack
	go get github.com/zetamatta/go-readline-ny
	go get github.com/mattn/go-colorable
	go get github.com/mdp/qrterminal
	mkdir -pv $(SRC)
	ln -sv $(PWD) $(SRC)
