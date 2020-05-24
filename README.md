<p align="center">
  <img src="/images/logo-small.png" alt="Logo of the Dimos project"
    title="Logo of the Dimos project" width="200" height="200">
</p>

# Dimos

Dimos (dee-mos, Δήμος: municipality, citizens) is a cryptocurrency based on a *Proof-of-Delegation* consensus algorithm.

The purpose of this cryptocurrency is to demonstrate that cryptocurrencies don't have to be profitable only for those who have the capital to invest in mining equipment or in units of crypto to stake. In the same way, it eliminates the need to maintain power-hungry miners (including ASIC) and creates a more egalitarian financial platform with fast transaction times.

The whitepaper is on its way.

## Getting Started

Install the dependencies and build the package:

``` sh
make install-deps && make
```

I don't currently use Go Modules for a few reason, first of which is that I couldn't find the specific versions of the dependencies that I wanted. It also nuked my environment when I tried using it. So until I'm forced to use it, I'll use Go Deps.

Now you can create the blockchain by initializing it.

``` sh
./bin/create-genesis
```

This should output something like this:

```
badger 2020/05/24 17:05:44 INFO: All 0 tables opened in 0s
2020/05/24 17:05:44 Merkle Root: a058b86554a94d7eaf219d18c4a87d455f6aa26ea0d066b5dc1b133b825c7b37 <nil>
2020/05/24 17:05:44 Genesis Hash: 106d0b419a37e7d388afb93f7b4ca5c7990d4700519fe5abd763b1477d56b2dad04ae0dc441ed0d028598ce1deb5aa193e0a2ae046f31fd1513377684df2470d <nil>
2020/05/24 17:05:44 Created genesis
```

Now you can run the `./bin/dimos-get-block` script and pass the genesis hash to it or just add the flag `-current`:

```
Usage of ./bin/dimos-get-block:
  -current
        Get the current block
  -hash string
        The hash of the block
```

#### Warning

This software is far from complete and improved every day.

