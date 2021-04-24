GOPATH := $(shell go env GOPATH)
SRC := $(GOPATH)/src/github.com/wisepythagoras/

all: validity wallet get-block genesis bg-service node helpers

download-deps:
	GO111MODULE=on go mod vendor

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

helpers:
	make -C cmd/test-block
	make -C cmd/dimos-test

test:
	go test ./crypto/
	go test ./proto/

coverage:
	./tools/gen_coverage.sh
