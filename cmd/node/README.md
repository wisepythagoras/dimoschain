# Node

In this folder you'll find the code that any node would run. Essentially, this is a server that accepts end-to-end encrypted connections from peers for handling transactions or updates. At the same time, nodes will need to keep track of the state of the network by repeating the messages that they receive to any other node they are aware of on the network. The most crucial functionality is how blocks are created and propagated throughout the network.

## Preparation

First install the dependencies.

``` sh
sudo apt install libssl-dev
# Or
sudo yum install openssl-devel
```

Then clone the repository of Themis and build and install it.

``` sh
git clone https://github.com/cossacklabs/themis
cd themis
make && sudo make install
```

The last step is to download the Go package.

``` sh
go get github.com/cossacklabs/themis/...
```

## Why Themis?

It's an easy way to use elliptic curve cryptography to get an end-to-end encrypted session featuring full forward secrecy in Go (and other languages).
