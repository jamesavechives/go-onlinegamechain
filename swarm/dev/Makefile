.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/onlinegamechain/go-onlinegamechain/cmd/swarm
	go build -o bin/geth     github.com/onlinegamechain/go-onlinegamechain/cmd/geth
	go build -o bin/bootnode github.com/onlinegamechain/go-onlinegamechain/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/onlinegamechain/go-onlinegamechain/swarm/...
