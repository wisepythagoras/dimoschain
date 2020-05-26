SRC := $(HOME)/go/src/github.com/wisepythagoras/

all:
	make -C cmd/dimos-bg-service
	make -C cmd/create-genesis
	make -C cmd/get-block
	make -C cmd/check-validity

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
	mkdir -pv $(SRC)
	ln -sv $(PWD) $(SRC)
